{
  description = "Flake for internet-chowkidar";

  inputs.nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";

  inputs.devshell.url = "github:numtide/devshell";
  inputs.devshell.inputs.nixpkgs.follows = "nixpkgs";

  inputs.flake-compat.url = "git+https://git.lix.systems/lix-project/flake-compat";
  inputs.flake-compat.flake = false;

  outputs =
    {
      self,
      nixpkgs,
      devshell,
      ...
    }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      eachSystem = nixpkgs.lib.genAttrs systems;
    in
    {
      devShells = eachSystem (
        system:
        let
          pkgs = import nixpkgs {
            inherit system;
            config.allowUnfree = true;
            overlays = [ devshell.overlays.default ];
          };
        in
        {
          default = pkgs.devshell.mkShell {
            bash = {
              interactive = "";
            };

            env = [
              {
                name = "DEVSHELL_NO_MOTD";
                value = 1;
              }
            ];

            packages = with pkgs; [
              git
              go-outline
              go
              gopls
              gotools
            ];
          };
        }
      );
      packages = eachSystem (
        system:
        let
          pkgs = import nixpkgs { inherit system; };
        in
        {
          default = pkgs.buildGoModule (finalAttrs: {
            pname = "chowkidar";
            version = "0.1.0";
            src = ./.;
            env.CGO_ENABLED = 1;
            
            ldflags = [
              "-X main.version=${finalAttrs.version}"
              "-X main.date=1970-01-01"
              "-X main.commit=${self.shortRev or "unknown"}"
            ];
            
            subPackages = [
              "cmd/chowkidar"
            ];
            
            vendorHash = "sha256-7KA64EdoDgwYqz2p72cEAouzhV9SInLmIF4JXd4fcuQ=";
            
            meta = {
              description = "A tool for detecting internet blocks by various ISPs";
              homepage = "https://inet.watch";
              license = pkgs.lib.licenses.mit;
            };
          });
        }
      );
    };
}

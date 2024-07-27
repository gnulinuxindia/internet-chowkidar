create name:
    @echo "Creating migration {{name}}..."
    cd migrations
    goose create {{name}}
    cd ..

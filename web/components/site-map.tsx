"use client";

import { useRef, useCallback, useState } from "react";
import Map, {
  Source,
  Layer,
  NavigationControl,
  type MapRef,
} from "react-map-gl/maplibre";
import type { ISP } from "@/lib/types";
import "maplibre-gl/dist/maplibre-gl.css";

interface SiteMapProps {
  isps: ISP[];
}

export function SiteMap({ isps }: SiteMapProps) {
  const mapRef = useRef<MapRef>(null);
  const [hoveredIsp, setHoveredIsp] = useState<ISP | null>(null);
  const [popupCoords, setPopupCoords] = useState<{ x: number; y: number } | null>(null);

  const geojsonData: GeoJSON.FeatureCollection = {
    type: "FeatureCollection",
    features: isps.map((isp) => ({
      type: "Feature",
      properties: {
        id: isp.id,
        name: isp.name,
        block_reports: isp.block_reports,
        unblock_reports: isp.unblock_reports,
      },
      geometry: {
        type: "Point",
        coordinates: [isp.longitude, isp.latitude],
      },
    })),
  };

  const maxBlocks = Math.max(...isps.map((isp) => isp.block_reports), 1);

  const onMouseMove = useCallback(
    (event: maplibregl.MapLayerMouseEvent) => {
      const features = event.features;
      if (features && features.length > 0) {
        const feature = features[0];
        const isp = isps.find((i) => i.id === feature.properties?.id);
        if (isp) {
          setHoveredIsp(isp);
          setPopupCoords({ x: event.point.x, y: event.point.y });
        }
      } else {
        setHoveredIsp(null);
        setPopupCoords(null);
      }
    },
    [isps]
  );

  const onMouseLeave = useCallback(() => {
    setHoveredIsp(null);
    setPopupCoords(null);
  }, []);

  return (
    <div className="relative map-container h-[500px] w-full">
      <Map
        ref={mapRef}
        initialViewState={{
          longitude: 78.9629,
          latitude: 20.5937,
          zoom: 4,
        }}
        style={{ width: "100%", height: "100%" }}
        mapStyle="https://basemaps.cartocdn.com/gl/dark-matter-gl-style/style.json"
        interactiveLayerIds={["isp-points", "isp-heatmap"]}
        onMouseMove={onMouseMove}
        onMouseLeave={onMouseLeave}
      >
        <NavigationControl position="top-right" />

        <Source id="isps" type="geojson" data={geojsonData}>
          {/* Heatmap layer - visible at lower zoom */}
          <Layer
            id="isp-heatmap"
            type="heatmap"
            maxzoom={9}
            paint={{
              "heatmap-weight": [
                "interpolate",
                ["linear"],
                ["get", "block_reports"],
                0,
                0,
                maxBlocks,
                1,
              ],
              "heatmap-intensity": [
                "interpolate",
                ["linear"],
                ["zoom"],
                0,
                1,
                9,
                3,
              ],
              "heatmap-color": [
                "interpolate",
                ["linear"],
                ["heatmap-density"],
                0,
                "rgba(0, 0, 0, 0)",
                0.2,
                "rgba(0, 200, 255, 0.3)",
                0.4,
                "rgba(0, 255, 200, 0.5)",
                0.6,
                "rgba(255, 200, 0, 0.7)",
                0.8,
                "rgba(255, 100, 50, 0.85)",
                1,
                "rgba(255, 50, 50, 1)",
              ],
              "heatmap-radius": [
                "interpolate",
                ["linear"],
                ["zoom"],
                0,
                20,
                9,
                50,
              ],
              "heatmap-opacity": [
                "interpolate",
                ["linear"],
                ["zoom"],
                7,
                1,
                9,
                0,
              ],
            }}
          />

          {/* Circle layer - visible at higher zoom */}
          <Layer
            id="isp-points"
            type="circle"
            minzoom={7}
            paint={{
              "circle-radius": [
                "interpolate",
                ["linear"],
                ["zoom"],
                7,
                ["interpolate", ["linear"], ["get", "block_reports"], 0, 4, maxBlocks, 12],
                16,
                ["interpolate", ["linear"], ["get", "block_reports"], 0, 8, maxBlocks, 24],
              ],
              "circle-color": [
                "interpolate",
                ["linear"],
                ["get", "block_reports"],
                0,
                "#00d4ff",
                maxBlocks * 0.3,
                "#00ff88",
                maxBlocks * 0.6,
                "#ffaa00",
                maxBlocks,
                "#ff5050",
              ],
              "circle-stroke-color": "#ffffff",
              "circle-stroke-width": 2,
              "circle-opacity": [
                "interpolate",
                ["linear"],
                ["zoom"],
                7,
                0,
                8,
                0.9,
              ],
            }}
          />
        </Source>
      </Map>

      {/* Custom tooltip */}
      {hoveredIsp && popupCoords && (
        <div
          className="absolute pointer-events-none z-10 px-4 py-3 bg-card border border-primary/30 rounded-lg shadow-lg"
          style={{
            left: popupCoords.x + 10,
            top: popupCoords.y + 10,
            transform: "translate(0, -50%)",
          }}
        >
          <p className="font-semibold text-foreground mb-1">{hoveredIsp.name}</p>
          <div className="flex gap-4 text-sm font-mono">
            <span className="text-[oklch(0.8_0.2_25)]">
              {hoveredIsp.block_reports} blocks
            </span>
            <span className="text-[oklch(0.8_0.18_145)]">
              {hoveredIsp.unblock_reports} unblocks
            </span>
          </div>
        </div>
      )}

      {/* Legend */}
      <div className="absolute bottom-4 left-4 bg-card/90 backdrop-blur-sm border border-border rounded-lg px-4 py-3">
        <p className="text-xs uppercase tracking-wider text-muted-foreground mb-2">
          Block Intensity
        </p>
        <div className="flex items-center gap-1">
          <div className="h-3 w-6 rounded-sm bg-[#00d4ff]" />
          <div className="h-3 w-6 rounded-sm bg-[#00ff88]" />
          <div className="h-3 w-6 rounded-sm bg-[#ffaa00]" />
          <div className="h-3 w-6 rounded-sm bg-[#ff5050]" />
        </div>
        <div className="flex justify-between text-xs text-muted-foreground mt-1">
          <span>Low</span>
          <span>High</span>
        </div>
      </div>
    </div>
  );
}

"use client";

import { Bar } from "react-chartjs-2";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
import { buildTimelineChartData } from "../../lib/chartData";
import type { Initiative } from "../../types/initiative";

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

interface TimelineBarProps {
  initiatives: Initiative[];
}

export function TimelineBar({ initiatives }: TimelineBarProps) {
  const data = buildTimelineChartData(initiatives);

  return (
    <div className="chart-container">
      <h2 className="chart-title">Announcements Over Time</h2>
      <Bar
        data={data}
        options={{
          responsive: true,
          plugins: {
            legend: { position: "bottom" },
          },
          scales: {
            x: { stacked: true },
            y: { stacked: true, ticks: { stepSize: 1 } },
          },
        }}
      />
    </div>
  );
}

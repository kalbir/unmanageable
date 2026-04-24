"use client";

import { Doughnut } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { buildStatusChartData } from "../../lib/chartData";
import type { Initiative } from "../../types/initiative";

ChartJS.register(ArcElement, Tooltip, Legend);

interface StatusDonutProps {
  initiatives: Initiative[];
}

export function StatusDonut({ initiatives }: StatusDonutProps) {
  const data = buildStatusChartData(initiatives);

  return (
    <div className="chart-container">
      <h2 className="chart-title">By Status</h2>
      <Doughnut
        data={data}
        options={{
          responsive: true,
          plugins: {
            legend: { position: "bottom" },
            tooltip: {
              callbacks: {
                label: (ctx) => `${ctx.label}: ${ctx.parsed}`,
              },
            },
          },
        }}
      />
    </div>
  );
}

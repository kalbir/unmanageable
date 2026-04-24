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
import { buildCategoryChartData } from "../../lib/chartData";
import type { Initiative } from "../../types/initiative";

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

interface CategoryBarProps {
  initiatives: Initiative[];
}

export function CategoryBar({ initiatives }: CategoryBarProps) {
  const data = buildCategoryChartData(initiatives);

  return (
    <div className="chart-container">
      <h2 className="chart-title">By Category</h2>
      <Bar
        data={data}
        options={{
          indexAxis: "y",
          responsive: true,
          plugins: {
            legend: { display: false },
          },
          scales: {
            x: {
              ticks: { stepSize: 1 },
            },
          },
        }}
      />
    </div>
  );
}

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
import { buildDepartmentChartData } from "../../lib/chartData";
import type { Initiative } from "../../types/initiative";

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

interface DepartmentBarProps {
  initiatives: Initiative[];
}

export function DepartmentBar({ initiatives }: DepartmentBarProps) {
  const data = buildDepartmentChartData(initiatives);

  return (
    <div className="chart-container">
      <h2 className="chart-title">By Department</h2>
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

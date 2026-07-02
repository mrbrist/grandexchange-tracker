import { useEffect, useRef } from "react";
import Chart, { Legend } from "chart.js/auto";
import "chartjs-adapter-date-fns";
import { type ItemPriceHistory } from "../api/getItem";

interface Props {
  history: ItemPriceHistory[];
}

function HistoryGraph({ history }: Props) {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const chartRef = useRef<Chart | null>(null);

  useEffect(() => {
    if (!canvasRef.current || history.length === 0) return;

    chartRef.current?.destroy();

    chartRef.current = new Chart(canvasRef.current, {
      type: "line",
      data: {
        datasets: [
          {
            label: "Average High Price",
            data: history.map((h) => ({
              x: h.priceTimestamp * 1000,
              y: h.avgHighPrice || null,
            })),
            tension: 0.1,
            spanGaps: true,
          },
          {
            label: "Average Low Price",
            data: history.map((h) => ({
              x: h.priceTimestamp * 1000,
              y: h.avgLowPrice || null,
            })),
            tension: 0.1,
            spanGaps: true,
          },
        ],
      },
      options: {
        responsive: true,
        plugins: {
          legend: {
            position: "bottom",
          },
        },
        interaction: {
          mode: "nearest",
          intersect: false,
        },
        scales: {
          x: {
            type: "time",
            time: {
              unit: "day",
              displayFormats: {
                day: "MMM d",
              },
            },
            title: {
              display: true,
              text: "Date",
            },
          },
          y: {
            title: {
              display: true,
              text: "Price",
            },
            beginAtZero: false,
          },
        },
      },
    });

    return () => {
      chartRef.current?.destroy();
    };
  }, [history]);

  return <canvas ref={canvasRef} />;
}

export default HistoryGraph;

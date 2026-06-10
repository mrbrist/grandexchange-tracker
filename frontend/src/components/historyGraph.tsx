import { useEffect, useRef } from "react";
import Chart from "chart.js/auto";
import { getRelativePosition } from "chart.js/helpers";
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
        labels: history.map((h) =>
          new Date(h.priceTimestamp * 1000).toLocaleDateString(),
        ),
        datasets: [
          {
            label: "Average High Price",
            data: history.map((h) =>
              h.avgHighPrice === 0 ? null : h.avgHighPrice,
            ),
            spanGaps: true,
            borderWidth: 2,
          },
          {
            label: "Average Low Price",
            data: history.map((h) =>
              h.avgLowPrice === 0 ? null : h.avgLowPrice,
            ),
            spanGaps: true,
            borderWidth: 2,
          },
        ],
      },
      options: {
        responsive: true,
        interaction: {
          mode: "nearest",
          intersect: false,
        },
        scales: {
          x: {
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
        onClick: (e) => {
          const chart = chartRef.current;
          if (!chart) return;

          const canvasPosition = getRelativePosition(e, chart);

          const dataX = chart.scales.x.getValueForPixel(canvasPosition.x);
          const dataY = chart.scales.y.getValueForPixel(canvasPosition.y);

          console.log("Clicked:", {
            dataIndex: dataX,
            value: dataY,
          });
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

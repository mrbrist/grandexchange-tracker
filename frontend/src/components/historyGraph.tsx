import { type ItemPriceHistory } from "../api/getItem";

function HistoryGraph({ history }: { history: ItemPriceHistory[] }) {
  return (
    <>
      {history.map((h) => (
        <p>{h.priceTimestamp}</p>
      ))}
    </>
  );
}

export default HistoryGraph;

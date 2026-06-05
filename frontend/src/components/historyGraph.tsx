import {
  useNavigate,
  useParams,
  type NavigateFunction,
} from "react-router-dom";
import { type ItemPriceHistory } from "../api/getItem";
import { useEffect, useState } from "react";

function HistoryGraph({ history }: { history: ItemPriceHistory[] }) {
  return <>{JSON.stringify(history)}</>;
}

export default HistoryGraph;

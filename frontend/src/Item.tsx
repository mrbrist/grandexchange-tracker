import {
  useNavigate,
  useParams,
  type NavigateFunction,
} from "react-router-dom";
import {
  getItem,
  type GetItemResponse,
  type ItemPriceHistory,
} from "./api/getItem";
import { useEffect, useState } from "react";

async function handleItemPrices(
  id: string,
  set: React.Dispatch<React.SetStateAction<GetItemResponse | undefined>>,
  navigate: NavigateFunction,
) {
  const item = await getItem(id);

  if (!item) {
    return;
  }

  if (!item?.history) {
    navigate("/", {
      state: {
        error: "There was no history found for that item",
      },
    });
  }

  set(item);
}

function Item() {
  const navigate = useNavigate();
  const { id } = useParams();

  const [item, setItem] = useState<GetItemResponse | undefined>();

  useEffect(() => {
    if (!id) {
      navigate("/");
      return;
    }

    handleItemPrices(id, setItem, navigate);
  }, [id, navigate]);

  return (
    <div>
      <h1>{item?.data.name ?? `Item ${id}`}</h1>

      {item?.data && (
        <div>
          <p>{item.data.examine}</p>
          <p>Value: {item.data.value.toLocaleString()}</p>
          <p>High Alch: {item.data.highalch.toLocaleString()}</p>
          <p>Low Alch: {item.data.lowalch.toLocaleString()}</p>
          <p>Members: {item.data.members ? "Yes" : "No"}</p>
          <p>GE Limit: {item.data.limit}</p>
        </div>
      )}

      <h2>Price History</h2>

      {item?.history.map((price: ItemPriceHistory) => (
        <div key={price.id}>
          <p>
            High:{" "}
            {price.avgHighPrice > 0
              ? price.avgHighPrice.toLocaleString()
              : "N/A"}
          </p>

          <p>
            Low:{" "}
            {price.avgLowPrice > 0 ? price.avgLowPrice.toLocaleString() : "N/A"}
          </p>

          <p>High Volume: {price.highVolume}</p>
          <p>Low Volume: {price.lowVolume}</p>

          <p>
            Timestamp: {new Date(price.priceTimestamp * 1000).toLocaleString()}
          </p>
        </div>
      ))}
    </div>
  );
}

export default Item;

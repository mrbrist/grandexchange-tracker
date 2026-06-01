import { useNavigate, useParams } from "react-router-dom";
import { getItem, type ItemPriceData } from "./api/getItem";
import { useEffect, useState } from "react";

async function handleItemPrices(
  id: string,
  set: React.Dispatch<React.SetStateAction<ItemPriceData[] | undefined>>,
) {
  const prices = await getItem(id);

  if (!prices) {
    return;
  }

  set(prices);
}

function Item() {
  const navigate = useNavigate();
  const { id } = useParams();

  const [prices, setPrices] = useState<ItemPriceData[] | undefined>();

  useEffect(() => {
    if (!id) {
      navigate("/");
      return;
    }

    handleItemPrices(id, setPrices);
  }, [id, navigate]);

  return (
    <div>
      <h1>Item {id}</h1>

      {prices?.map((price) => (
        <div key={price.ID}>
          <p>High: {price.AvgHighPrice}</p>
          <p>Low: {price.AvgLowPrice}</p>
          <p>Timestamp: {price.PriceTimestamp}</p>
        </div>
      ))}
    </div>
  );
}

export default Item;

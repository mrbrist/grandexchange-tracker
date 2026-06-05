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
    // <h1>{item?.data.name ?? `Item ${id}`}</h1>

    // {item?.data && (
    //   <div>
    //     <p>{item.data.examine}</p>
    //     <p>Value: {item.data.value.toLocaleString()}</p>
    //     <p>High Alch: {item.data.highalch.toLocaleString()}</p>
    //     <p>Low Alch: {item.data.lowalch.toLocaleString()}</p>
    //     <p>Members: {item.data.members ? "Yes" : "No"}</p>
    //     <p>GE Limit: {item.data.limit}</p>
    //   </div>
    // )}
    <>
      <div className="min-h-screen bg-zinc-950 text-white flex flex-col items-center pt-10 px-4">
        <a href="/" className="text-5xl font-bold mb-10 tracking-tight">
          {item?.data.name ?? `Item ${id}`}
        </a>

        <div className="w-full max-w-2xl">
          {/* <div className="flex items-center bg-zinc-900 border border-zinc-700 rounded-2xl overflow-hidden shadow-lg"></div> */}
        </div>
      </div>
    </>
  );
}

export default Item;

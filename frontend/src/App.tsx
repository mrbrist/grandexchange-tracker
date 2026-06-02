import { useEffect, useState } from "react";
import { getList, type ListData } from "./api/getList";

async function handleItemList(
  set: React.Dispatch<React.SetStateAction<ListData[] | undefined>>,
) {
  const list = await getList();

  if (!list) {
    return;
  }

  set(list);
}

function App() {
  const [list, setList] = useState<ListData[] | undefined>();
  const [query, setQuery] = useState("");

  useEffect(() => {
    handleItemList(setList);
  }, []);

  console.log(list);

  return (
    <>
      <div className="min-h-screen bg-zinc-950 text-white flex flex-col items-center pt-20 px-4">
        <h1 className="text-5xl font-bold mb-10 tracking-tight">
          Grand Exchange Tracker
        </h1>

        <div className="w-full max-w-2xl">
          <div className="flex items-center bg-zinc-900 border border-zinc-700 rounded-2xl overflow-hidden shadow-lg">
            <input
              type="text"
              placeholder="Search for an item..."
              className="flex-1 px-6 py-5 bg-transparent text-lg outline-none placeholder:text-zinc-500"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
            />
          </div>
          {query}
        </div>
      </div>
    </>
  );
}

export default App;

import { useEffect, useState } from "react";
import { getList, type ListData } from "./api/getList";
import "./App.css";

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

  useEffect(() => {
    handleItemList(setList);
  });

  return (
    <>
      {list?.map((i) => (
        <p>{i.name}</p>
      ))}
    </>
  );
}

export default App;

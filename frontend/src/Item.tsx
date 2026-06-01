import { useParams } from "react-router-dom";
function Item() {
  const { id } = useParams();
  return <>{id}</>;
}

export default Item;

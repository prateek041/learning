import Card from "./card";
import { Todo } from "../types/types";
import { IListItemsProps } from "../types/types";

const ListItems = ({ todos }: IListItemsProps) => {
  return (
    <div className="flex mx-5 gap-3 justify-center w-full flex-wrap">
      {todos.map((todo: Todo) => (
        <Card title={todo.title} description={todo.description} />
      ))}
    </div>
  );
};

export default ListItems;

import Card from "./card";

export interface Todo {
  title: string;
  description: string;
}

const ListItems = ({ todos }: { todos: Todo[] }) => {
  return (
    <div className="flex mx-5 gap-3 justify-center w-full flex-wrap">
      {todos.map((todo: Todo) => (
        <Card title={todo.title} description={todo.description} />
      ))}
    </div>
  );
};

export default ListItems;

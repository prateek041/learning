import ListItems from "./components/listitems";
import { Todo } from "./types/types";

const sampleTodos: Todo[] = [
  {
    title: "hello",
    description: "first todo",
  },
  {
    title: "world",
    description: "second todo",
  },
  {
    title: "this is new",
    description: "third todo",
  },
  {
    title: "one more",
    description: "final one",
  },
  {
    title: "hello",
    description: "first todo",
  },
  {
    title: "world",
    description: "second todo",
  },
  {
    title: "this is new",
    description: "third todo",
  },
  {
    title: "one more",
    description: "final one",
  },
];

const App = () => {
  return (
    <div className="h-screen flex flex-col">
      <h1 className="text-6xl font-bold mb-10 mx-5 mt-5">Your TodoList</h1>
      <div className="flex justify-center self-center w-3/4 h-full">
        <ListItems todos={sampleTodos} />
      </div>
    </div>
  );
};

export default App;

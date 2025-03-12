import { ICardProps } from "../types/types";

const Card = ({ title, description }: ICardProps) => {
  const testOnClick = () => {
    console.log("I was clicked");
  };

  return (
    <div className="border w-1/3 h-1/3 flex flex-col justify-between px-5 py-3">
      <div className="">
        <h1 className="text-3xl">{title}</h1>
        <p className="text-xl">{description}</p>
      </div>
      <div>
        <button onClick={testOnClick}>Remove</button>
      </div>
    </div>
  );
};

export default Card;

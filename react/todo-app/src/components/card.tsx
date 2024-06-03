const Card = ({
  title,
  description,
}: {
  title: string;
  description: string;
}) => {
  return (
    <div className="border w-1/3 h-1/3 px-5 py-3">
      <h1 className="text-3xl">{title}</h1>
      <p className="text-xl">{description}</p>
    </div>
  );
};

export default Card;

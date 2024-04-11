import "@/index.css";

const Message = ({ message }) => {
  const time = message.ts.match(/\d\d:\d\d/);
  return (
    <div className="flex w-max items-start p-5 bg-gray-50 border border-gray-300 rounded-lg">
      <div className="flex flex-col w-full max-w-[320px] leading-1.5">
        <div className="flex items-center space-x-2 rtl:space-x-reverse">
          <span className="text-sm font-semibold text-gray-900">
            {message.sender_name}
          </span>
          <span className="text-sm font-normal text-gray-500">{time}</span>
        </div>
        <p className="text-sm font-normal py-2 text-gray-900">{message.data}</p>
        <span className="text-sm font-normal text-gray-500">Delivered</span>
      </div>
    </div>
  );
};

export default Message;

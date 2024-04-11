import "@/index.css";

const Message = ({ message }) => {
  return (
    <div className="flex items-start p-5">
      <div className="flex flex-col w-full max-w-[320px] leading-1.5">
        <div className="flex items-center space-x-2 rtl:space-x-reverse">
          <span className="text-sm font-semibold text-gray-900">
            {message.sender_name}
          </span>
          <span className="text-sm font-normal text-gray-500">
            {message.ts}
          </span>
        </div>
        <p className="text-sm font-normal py-2 text-gray-900">{message.data}</p>
        <span className="text-sm font-normal text-gray-500">Delivered</span>
      </div>
    </div>
  );
};

export default Message;

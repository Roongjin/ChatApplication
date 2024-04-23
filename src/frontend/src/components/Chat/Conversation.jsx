import "@/index.css";
import { useState } from "react";

const Message = ({ message }) => {
  const time = message.ts.match(/\d\d:\d\d/);
  return (
    <div className="flex w-full max-w-full items-start p-5 bg-gray-50 border border-gray-300 rounded-lg">
      <div className="flex flex-col w-full max-w-[320px] leading-1.5">
        <div className="flex items-center space-x-2 rtl:space-x-reverse">
          <span className="text-sm font-semibold text-gray-900">
            {message.sender_name}
          </span>
          <span className="text-sm font-normal text-gray-500">{time}</span>
        </div>
        <p className="text-sm font-normal py-2 text-gray-900 break-words">
          {message.text}
        </p>
        <span className="text-sm font-normal text-gray-500">Delivered</span>
      </div>
    </div>
  );
};

const InputBox = ({ userId, roomId, sendJsonMessage }) => {
  const [input, setInput] = useState("");

  const handleTyping = (event) => {
    event.preventDefault();
    setInput(event.target.value);
    sendJsonMessage({
      text: "1",
      type: "typing",
      sender: userId,
      room: roomId,
    });
  };

  const sendMessage = (event) => {
    event.preventDefault();
    sendJsonMessage({
      text: input,
      type: "message",
      sender: userId,
      room: roomId,
    });
    setInput("");
  };

  return (
    <form onSubmit={sendMessage} className="border-gray-100 border rounded-lg">
      <div className="flex items-center px-3 py-2 rounded-lg bg-gray-50">
        <input
          id="chat"
          rows="1"
          className="block mx-4 p-2.5 w-full text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
          value={input}
          onChange={handleTyping}
          placeholder="Your message..."
        ></input>
        <button
          type="submit"
          className="inline-flex justify-center p-2 text-blue-600 rounded-full cursor-pointer hover:bg-blue-100"
        >
          <svg
            className="w-5 h-5 rotate-90 rtl:-rotate-90"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="currentColor"
            viewBox="0 0 18 20"
          >
            <path d="m17.914 18.594-8-18a1 1 0 0 0-1.828 0l-8 18a1 1 0 0 0 1.157 1.376L8 18.281V9a1 1 0 0 1 2 0v9.281l6.758 1.689a1 1 0 0 0 1.156-1.376Z" />
          </svg>
          <span className="sr-only">Send message</span>
        </button>
      </div>
    </form>
  );
};

export { Message, InputBox };

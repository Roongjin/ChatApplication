import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import apiClient from "@/libs/apiClient";
import "@/index.css";
import Message from "@/components/Chat/Room";
import OnlineUsers from "@/components/Chat/Sidebar";
import useWebSocket, { ReadyState } from "react-use-websocket";

const Chat = ({ userId }) => {
  const [messages, setMessages] = useState([]);
  const [onlineUsers, setOnlineUsers] = useState([]);
  const navigate = useNavigate();
  const WS_URL = `ws://localhost:8080/chat/ws/${userId}`;
  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(
    WS_URL,
    {
      share: false,
      shouldReconnect: () => true,
    },
  );

  //On boot
  useEffect(() => {
    if (!userId) {
      navigate("/");
    }

    async function fetchOnlineUsers() {
      const { data } = await apiClient
        .get("/chat/online-users")
        .then((resp) => resp.data);
      console.log("data", data);
      if (data) {
        Array.from(data).forEach((e) => {
          console.log(e);
        });
      }
      setOnlineUsers(data);
    }
    fetchOnlineUsers();
  }, []);

  //Chat
  useEffect(() => {
    console.log("Connection state changed");
    if (readyState === ReadyState.OPEN) {
      sendJsonMessage({
        data: `${userId} has connected from front end!`,
        type: "message",
        sender: `${userId}`,
        room: "28e5b39a-e00f-4f10-9975-bc25ff0afacb",
      });
    }
  }, [readyState]);

  useEffect(() => {
    if (!lastJsonMessage) {
      return;
    }
    console.log(lastJsonMessage);
    console.log(messages);
    if (lastJsonMessage.type === "message") {
      setMessages(messages.concat(lastJsonMessage));
    }
  }, [lastJsonMessage]);

  return (
    <>
      <div className="grid grid-cols-3 gap-4">
        <div className="max-w-full">
          {messages.map((msg) => (
            <Message key={msg.id} message={msg} />
          ))}
        </div>
        <div>
          <OnlineUsers onlineUsers={onlineUsers} selfId={userId} />
        </div>
      </div>
    </>
  );
};

export default Chat;

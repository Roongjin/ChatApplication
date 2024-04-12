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
  const WS_URL = `ws://${import.meta.env.VITE_IPADDR}:8080/chat/ws/${userId}`;
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
      if (data) {
        setOnlineUsers(data);
      }
    }
    fetchOnlineUsers();
  }, []);

  //Chat
  useEffect(() => {
    console.log("Connection state changed");
    if (readyState === ReadyState.OPEN) {
      console.log("Connection ready");
    }
  }, [readyState]);

  useEffect(() => {
    if (!lastJsonMessage) {
      return;
    }
    console.log(lastJsonMessage);
    console.log(messages);
    console.log(onlineUsers);

    if (lastJsonMessage.type === "message") {
      setMessages(messages.concat(lastJsonMessage));
      return;
    }

    if (lastJsonMessage.type === "presence") {
      if (lastJsonMessage.text === "1") {
        setOnlineUsers(
          onlineUsers.concat({
            id: lastJsonMessage.sender,
            is_online: true,
            username: lastJsonMessage.sender_name,
          }),
        );
        return;
      }
      setOnlineUsers(
        onlineUsers.filter(
          (user) => user.username !== lastJsonMessage.sender_name,
        ),
      );
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

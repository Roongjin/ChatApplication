import { useState, useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import apiClient from "@/libs/apiClient";
import "@/index.css";
import { Message, InputBox } from "@/components/Chat/Conversation";
import { AllUsers, NewRoomButton } from "@/components/Chat/Sidebar";
import useWebSocket, { ReadyState } from "react-use-websocket";
import RoomList from "./Room";

const Chat = ({ userId }) => {
  const [messages, setMessages] = useState([]);
  const [allUsers, setAllUsers] = useState([]);
  const [currentRoomId, setCurrentRoomId] = useState("");
  const [existedRooms, setExistedRooms] = useState([]);
  const [wsUrl, setWsUrl] = useState(
    `ws://${import.meta.env.VITE_IPADDR}:8080/chat/ws/${userId}`,
  );
  const [bcstRoomId, setBcstRoomId] = useState("");
  const chatContainerRef = useRef(null);
  const navigate = useNavigate();

  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(wsUrl, {
    share: false,
    shouldReconnect: () => true,
  });

  useEffect(() => {
    console.log("wsurl changed");
    setWsUrl(`ws://${import.meta.env.VITE_IPADDR}:8080/chat/ws/${userId}`);
  }, [wsUrl]);

  //On boot
  useEffect(() => {
    if (!userId) {
      navigate("/");
    }

    async function fetchAllUsers() {
      const { data } = await apiClient
        .get("/chat/all-users")
        .then((resp) => resp.data);
      if (data) {
        setAllUsers(data);
      }
    }

    async function fetchExistedRooms() {
      const { data } = await apiClient
        .get(`/chat/${userId}`)
        .then((resp) => resp.data);
      if (data) {
        setExistedRooms(data);
      }
    }

    async function fetchBcstRoomId() {
      const { data } = await apiClient
        .get("/chat/broadcast-id")
        .then((resp) => resp.data);
      if (data) {
        setBcstRoomId(data);
      }
    }

    fetchAllUsers();
    fetchExistedRooms();
    fetchBcstRoomId();
  }, [wsUrl]);

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
    console.log("latest message: ", lastJsonMessage);
    console.log("all messages: ", messages);
    console.log("all user instances: ", allUsers);

    if (lastJsonMessage.type === "message") {
      setMessages(messages.concat(lastJsonMessage));
      return;
    }

    if (lastJsonMessage.type === "presence") {
      if (lastJsonMessage.text === "1") {
        setAllUsers(
          allUsers.some((user) => user.id === lastJsonMessage.sender)
            ? allUsers.map((user) =>
                user.id === lastJsonMessage.sender
                  ? { ...user, is_online: true }
                  : user,
              )
            : allUsers.concat({
                id: lastJsonMessage.sender,
                is_online: true,
                username: lastJsonMessage.sender_name,
              }),
        );
        return;
      }
      setAllUsers(
        allUsers.map((user) =>
          user.id === lastJsonMessage.sender
            ? { ...user, is_online: false }
            : user,
        ),
      );
    }

    if (lastJsonMessage.type === "reset") {
      setWsUrl("");
    }
  }, [lastJsonMessage]);

  useEffect(() => {
    chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
  }, [messages]);

  return (
    <>
      <div className="grid grid-cols-3 gap-4">
        <div>
          <RoomList
            existedRooms={existedRooms}
            setCurrentRoomId={setCurrentRoomId}
            setMessages={setMessages}
          />
        </div>
        <div className="max-w-full relative flex flex-col border border-gray-300 rounded-lg">
          <div
            className="relative overflow-y-scroll no-scrollbar max-h-[90vh] h-[90vh]"
            ref={chatContainerRef}
          >
            {messages.map((msg) => (
              <Message key={msg.id} message={msg} />
            ))}
          </div>
          <div className="relative m-2">
            <InputBox
              userId={userId}
              roomId={currentRoomId}
              sendJsonMessage={sendJsonMessage}
            />
          </div>
        </div>
        <div>
          <div>
            <AllUsers allUsers={allUsers} selfId={userId} />
          </div>
          <div>
            <NewRoomButton
              selfId={userId}
              existedRooms={existedRooms}
              setExistedRooms={setExistedRooms}
              setCurrentRoomId={setCurrentRoomId}
              setMessages={setMessages}
              setWsUrl={setWsUrl}
              sendJsonMessage={sendJsonMessage}
              bcstRoomId={bcstRoomId}
            />
          </div>
        </div>
      </div>
    </>
  );
};

export default Chat;

import { useState, useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import apiClient from "@/libs/apiClient";
import "@/index.css";
import { Message, InputBox } from "@/components/Chat/Conversation";
import RoomList from "@/components/Chat/Room";
import { AllUsers, NewRoomButton } from "@/components/Chat/Sidebar";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { Button, Modal } from "flowbite-react";

const Chat = ({ userId }) => {
  const [currentTypingUsers, setCurrentTypingUsers] = useState([]);
  const [messages, setMessages] = useState([]);
  const [allUsers, setAllUsers] = useState([]);
  const [currentRoomId, setCurrentRoomId] = useState("");
  const [existedRooms, setExistedRooms] = useState([]);
  const [openModal, setOpenModal] = useState(false);
  const [roomMembers, setRoomMembers] = useState([]);
  const [roomMemberInput, setRoomMemberInput] = useState("");
  const [wsUrl, setWsUrl] = useState(
    `ws://${import.meta.env.VITE_IPADDR}:8080/chat/ws/${userId}`,
  );
  const [bcstRoomId, setBcstRoomId] = useState("");
  const chatContainerRef = useRef(null);
  const timeoutRef = useRef({});
  const navigate = useNavigate();

  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(wsUrl, {
    share: false,
    shouldReconnect: () => true,
  });

  async function fetchExistedRooms() {
    const { data } = await apiClient
      .get(`/chat/${userId}`)
      .then((resp) => resp.data);
    if (data) {
      setExistedRooms(data);
    }
  }

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
    console.log("latest message: ", lastJsonMessage);
    console.log("all messages: ", messages);
    console.log("all user instances: ", allUsers);

    if (lastJsonMessage.type === "message") {
      fetchExistedRooms();
      if (lastJsonMessage.room === currentRoomId) {
        setMessages(messages.concat(lastJsonMessage));
      }
      return;
    }

    if (
      lastJsonMessage.type === "typing" &&
      lastJsonMessage.room === currentRoomId &&
      lastJsonMessage.sender !== userId
    ) {
      if (
        currentTypingUsers.some((user) => user.id === lastJsonMessage.sender)
      ) {
        clearTimeout(timeoutRef.current[lastJsonMessage.sender]);
        timeoutRef.current[lastJsonMessage.sender] = setTimeout(
          () =>
            setCurrentTypingUsers(
              currentTypingUsers.filter(
                (user) => user.id !== lastJsonMessage.sender,
              ),
            ),
          1000,
        );
      } else {
        setCurrentTypingUsers(
          currentTypingUsers.concat({
            id: lastJsonMessage.sender,
            username: lastJsonMessage.sender_name,
          }),
        );
        timeoutRef.current[lastJsonMessage.sender] = setTimeout(
          () =>
            setCurrentTypingUsers(
              currentTypingUsers.filter(
                (user) => user.id !== lastJsonMessage.sender,
              ),
            ),
          1000,
        );
      }

      console.log("Typing: ", currentTypingUsers);
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
    setCurrentTypingUsers([]);
  }, [currentRoomId]);

  useEffect(() => {
    chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
  }, [messages]);

  const closeModal = () => {
    setOpenModal(false);
    setRoomMembers([]);
  };

  const HandleConfirmNewRoom = async () => {
    setOpenModal(false);
    setRoomMembers([]);
    const { data } = await apiClient
      .post(`/chat/new-room/${userId}`, {
        room_members_name: roomMembers,
      })
      .then((resp) => resp.data);

    if (!existedRooms.some((room) => room.id === data.id)) {
      setExistedRooms(existedRooms.concat(data));
    }

    setCurrentRoomId(data.id);

    const messages = await apiClient
      .get(`/chat/conv/${data.id}`)
      .then((resp) => resp.data);
    if (messages) {
      setMessages(messages.data);
    }

    setWsUrl("");
    sendJsonMessage({
      text: "reset",
      type: "reset",
      sender: userId,
      room: bcstRoomId,
    });
  };

  //TODO: set scrollable for room list
  return (
    <>
      <Modal
        dismissible
        show={openModal}
        onClose={closeModal}
        className="w-1/3 m-auto flex justify-center items-center"
      >
        <div className="bg-blue-100 rounded-lg border border-gray-300 w-full">
          <Modal.Header className="bg-blue-200 border rounded-t-lg">
            <span>Enter room members</span>
          </Modal.Header>
          <Modal.Body>
            <div className="space-y-6 m-2">
              {roomMembers.map((member) => (
                <span
                  key={member}
                  className="border border-gray-300 rounded-lg bg-blue-200 m-2"
                >
                  {member}
                </span>
              ))}
            </div>
            <form
              onSubmit={(event) => {
                event.preventDefault();
                setRoomMembers(roomMembers.concat(roomMemberInput));
                setRoomMemberInput("");
              }}
            >
              <input
                className="block mx-2 m-3 p-3 w-1/2 text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
                value={roomMemberInput}
                onChange={(e) => setRoomMemberInput(e.target.value)}
                placeholder="Username"
              ></input>
            </form>
            <Button
              onClick={HandleConfirmNewRoom}
              className="bg-blue-200 m-2 p-2"
            >
              Confirm
            </Button>
          </Modal.Body>
        </div>
      </Modal>

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
            {currentTypingUsers.map((user) => (
              <p key={user.id}>{user.username} is typing...</p>
            ))}
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
            <NewRoomButton setOpenModal={setOpenModal} />
          </div>
        </div>
      </div>
    </>
  );
};

export default Chat;

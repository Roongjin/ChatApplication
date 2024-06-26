import "@/index.css";
import apiClient from "@/libs/apiClient";
import { Button, Modal } from "flowbite-react";
import { useState } from "react";

const AllUsers = ({ allUsers, selfId }) => {
  if (!allUsers || allUsers.length === 0) {
    return (
      <p className="m-5 p-1 w-max max-w-fit text-sm font-medium text-gray-900 bg-gray-50 border border-gray-300 rounded-lg">
        No users currently existed
      </p>
    );
  }

  return (
    <>
      <p className="m-5 p-1 w-max max-w-fit text-sm font-medium text-gray-900 bg-gray-50 border border-gray-300 rounded-lg">
        Users
      </p>
      {allUsers.map((user) =>
        user.id === selfId ? null : (
          <p key={user.id} className="p-5">
            <span className="flex items-center text-sm font-medium text-gray-900 me-3">
              {user.is_online ? (
                <span className="flex w-2.5 h-2.5 bg-green-400 rounded-full me-1.5 flex-shrink-0"></span>
              ) : (
                <span className="flex w-2.5 h-2.5 bg-gray-600 rounded-full me-1.5 flex-shrink-0"></span>
              )}
              {user.username}
            </span>
          </p>
        ),
      )}
    </>
  );
};

const NewRoomButton = ({
  selfId,
  existedRooms,
  setExistedRooms,
  setCurrentRoomId,
  setMessages,
  setWsUrl,
  sendJsonMessage,
  bcstRoomId,
}) => {
  const [openModal, setOpenModal] = useState(false);
  const [roomMembers, setRoomMembers] = useState([]);
  const [input, setInput] = useState("");

  const HandleConfirmNewRoom = async () => {
    setOpenModal(false);
    setRoomMembers([]);
    const { data } = await apiClient
      .post(`/chat/new-room/${selfId}`, {
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
      sender: selfId,
      room: bcstRoomId,
    });
  };

  const closeModal = () => {
    setOpenModal(false);
    setRoomMembers([]);
  };

  return (
    <>
      <Button onClick={() => setOpenModal(true)} className="bg-gray-600 p-2">
        New Room
      </Button>
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
                setRoomMembers(roomMembers.concat(input));
                setInput("");
              }}
            >
              <input
                className="block mx-2 m-3 p-3 w-1/2 text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
                value={input}
                onChange={(e) => setInput(e.target.value)}
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
    </>
  );
};

export { AllUsers, NewRoomButton };

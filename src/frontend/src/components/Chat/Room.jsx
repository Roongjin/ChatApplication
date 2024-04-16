import apiClient from "@/libs/apiClient";

const RoomList = ({ existedRooms, setCurrentRoomId, setMessages }) => {
  return (
    <>
      {existedRooms.map((room) => (
        <Room
          key={room.id}
          room={room}
          handleRoomClick={async () => {
            setCurrentRoomId(room.id);

            const { data } = await apiClient
              .get(`/chat/conv/${room.id}`)
              .then((resp) => resp.data);
            if (data) {
              setMessages(data);
            }
          }}
        />
      ))}
    </>
  );
};

const Room = ({ room, handleRoomClick }) => {
  return (
    <div
      onClick={handleRoomClick}
      className="flex w-full max-w-full items-center p-5 bg-gray-50 border border-gray-300 rounded-lg m-2"
    >
      {room.members.map((user) => (
        <span
          key={user.id}
          className="bg-green-100 border border-gray-300 rounded-lg p-1 m-1"
        >
          {user.username}
        </span>
      ))}
    </div>
  );
};

export default RoomList;

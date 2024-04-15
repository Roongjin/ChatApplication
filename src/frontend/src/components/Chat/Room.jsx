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
              console.log(data);
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
      className="flex w-max max-w-fit items-start p-5 bg-gray-50 border border-gray-300 rounded-lg"
    >
      {room.id}
    </div>
  );
};

export default RoomList;

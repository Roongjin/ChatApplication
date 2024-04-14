import "@/index.css";

const OnlineUsers = ({ onlineUsers, selfId }) => {
  if (!onlineUsers || onlineUsers.length === 0) {
    return (
      <p className="m-5 p-1 w-max max-w-fit text-sm font-medium text-gray-900 bg-gray-50 border border-gray-300 rounded-lg">
        No users currently online
      </p>
    );
  }

  return (
    <>
      <p className="m-5 p-1 w-max max-w-fit text-sm font-medium text-gray-900 bg-gray-50 border border-gray-300 rounded-lg">
        Online Users
      </p>
      {onlineUsers.map((user) =>
        user.id === selfId ? null : (
          <p key={user.id} className="p-5">
            <span className="flex items-center text-sm font-medium text-gray-900 me-3">
              <span className="flex w-2.5 h-2.5 bg-blue-600 rounded-full me-1.5 flex-shrink-0"></span>
              {user.username}
            </span>
          </p>
        ),
      )}
    </>
  );
};

export default OnlineUsers;

import "@/index.css";

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

export default AllUsers;

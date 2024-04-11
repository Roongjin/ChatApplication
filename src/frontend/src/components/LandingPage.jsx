import { useState } from "react";
import apiClient from "../libs/apiClient";
import { useNavigate } from "react-router-dom";
import "../index.css";

const LandingPage = ({ setUserId }) => {
  const navigate = useNavigate();
  const [name, setName] = useState("");

  const register = async () => {
    try {
      const { data } = await apiClient
        .post(`/authen/${name}`)
        .then((resp) => resp.data);
      setUserId(data.id);
      navigate("/chat");
    } catch (err) {
      console.log(err);
      throw err;
    }
  };

  return (
    <div className="max-w-sm mx-auto">
      <div className="m-5">
        <input
          onChange={(event) => setName(event.target.value)}
          className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 mb-5"
          placeholder="Please enter your name"
        />
        <button
          type="submit"
          onClick={register}
          className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        >
          Confirm
        </button>
      </div>
    </div>
  );
};

export default LandingPage;

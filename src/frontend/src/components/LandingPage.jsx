import { useState } from "react";
import apiClient from "../libs/apiClient";
import { useNavigate } from "react-router-dom";
import "../index.css";

const LandingPage = ({ setUserId }) => {
  const navigate = useNavigate();
  const [name, setName] = useState("");

  const register = async () => {
    try {
      const { data } = await apiClient.post(`/authen/${name}`);
      setUserId(data.data.id);
      navigate("/chat");
    } catch (err) {
      console.log(err);
      throw err;
    }
  };

  return (
    <>
      <h3 className="text-sky-600 w-1/2 rounded">Please enter your name</h3>
      <input
        onChange={(event) => setName(event.target.value)}
        className="bg-black border-2 border-white text-white rounded w-1/2"
      />
      <button
        type="submit"
        onClick={register}
        className="bg-sky-300 rounded w-1/2"
      >
        Confirm
      </button>
    </>
  );
};

export default LandingPage;

import { useState } from "react";
import apiClient from "../libs/apiClient";
import { useNavigate } from "react-router-dom";

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
      <h3>Please enter your name</h3>
      <input onChange={(event) => setName(event.target.value)} />
      <button type="submit" onClick={register}>
        Confirm
      </button>
    </>
  );
};

export default LandingPage;

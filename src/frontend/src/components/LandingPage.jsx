import { useState } from "react";
import apiClient from "../libs/apiClient";

const LandingPage = () => {
  const [name, setName] = useState("");

  const register = async () => {
    try {
      const { data } = await apiClient.post(`/authen/${name}`);
      return data;
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

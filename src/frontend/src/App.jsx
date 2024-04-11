import LandingPage from "@/components/LandingPage/LandingPage";
import Chat from "@/components/Chat/Chat";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { useState } from "react";

function App() {
  const [userId, setUserId] = useState("");
  return (
    <Router>
      <Routes>
        <Route path="/chat" element={<Chat userId={userId} />} />
        <Route path="/" element={<LandingPage setUserId={setUserId} />} />
      </Routes>
    </Router>
  );
}

export default App;

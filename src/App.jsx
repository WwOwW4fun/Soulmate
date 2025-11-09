import { useState, useEffect, useRef } from "react";
import "./App.css";

function App() {
  const [messages, setMessages] = useState([
    { sender: "bot", text: "Hi, I'm Soulmate 💜 How are you feeling today?" },
  ]);
  const [input, setInput] = useState("");
  const [isTyping, setIsTyping] = useState(false);
  const chatEndRef = useRef(null);

  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages, isTyping]);

  const handleSend = async () => {
    if (!input.trim()) return;

    const newMessage = { sender: "user", text: input };
    setMessages((prev) => [...prev, newMessage]);
    setInput("");
    setIsTyping(true);

  const res = await fetch("http://localhost:8080/chatguided", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ message: input }),
  });

   const data = await res.json();

  if (data.reply) {
    setMessages((prev) => [...prev, { sender: "bot", text: data.reply }]);
  }

  /*if (data.healingPlan) {
    setMessages((prev) => [...prev, { sender: "bot", text: data.healingPlan }]);
  }*/

  setIsTyping(false);
  };

  return (
    <div className="app">
      <h1 className="title">💜 SoulMate</h1>

      <div className="chat-container">
        <div className="chat-box">
          {messages.map((msg, i) => (
            <div
              key={i}
              className={`message ${msg.sender === "user" ? "user" : "bot"}`}
            >
              {msg.text}
            </div>
          ))}

          {isTyping && (
            <div className="typing">
              <span></span>
              <span></span>
              <span></span>
            </div>
          )}

          <div ref={chatEndRef} />
        </div>

        <div className="input-container">
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleSend()}
            placeholder="Let share your feeling"
          />
          <button onClick={handleSend}>Send</button>
        </div>
      </div>

      <footer>
      © 2025 Soulmate - Created by An Nguyen, Khoa Danh, Hiep Nguyen, Phuc Ho
      </footer>
    </div>
  );
}

export default App;

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

  const handleSend = () => {
    if (!input.trim()) return;

    const newMessage = { sender: "user", text: input };
    setMessages((prev) => [...prev, newMessage]);
    setInput("");

    setIsTyping(true);
    setTimeout(() => {
      setIsTyping(false);
      setMessages((prev) => [
        ...prev,
        {
          sender: "bot",
          text:
            "That sounds really meaningful. Can you share a bit more about how that makes you feel?",
        },
      ]);
    }, 2000);
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
      © 2025 Soulmate - Created by An Nguyen, Khoa Danh, Hiep Nguyen, Minh Trinh
      </footer>
    </div>
  );
}

export default App;

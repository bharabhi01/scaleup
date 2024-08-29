import React from "react";

const ServerHealth = ({ server }) => {
  const { name, address, alive } = server;

  return (
    <div>
      <h2>{name}</h2>
      <p>IP address: {address}</p>
      <p>Status: {alive ? "Healthy" : "Unhealty"}</p>
    </div>
  );
};

export default ServerHealth;

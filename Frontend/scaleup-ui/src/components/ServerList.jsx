import React from "react";
import ServerHealth from "./ServerHealth";

const ServerList = ({ servers }) => {
  if (ServerList.length === 0) {
    return <div>No servers available</div>;
  }

  return (
    <div>
      {servers.map((server, index) => (
        <ServerHealth key={index} server={server} />
      ))}
    </div>
  );
};

export default ServerList;

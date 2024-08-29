import React from "react";
import useFetchServers from "../hooks/useFetchServer";
import ServerList from "../components/ServerList";

const Dashboard = () => {
  const { servers, loading, error } = useFetchServers();

  if (loading) {
    return <div>Loading Server data...</div>;
  }

  if (error) {
    return <div>Error loading Server data : {error}</div>;
  }

  return (
    <div>
      <h1>Server Health Dashboard</h1>
      <ServerList servers={servers} />
    </div>
  );
};

export default Dashboard;

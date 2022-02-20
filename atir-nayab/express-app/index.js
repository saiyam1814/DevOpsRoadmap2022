const express = require("express");
const mongoose = require("mongoose");
const {
  MONGO_USER,
  MONGO_PASSWORD,
  MONGO_IP,
  MONGO_PORT,
} = require("./config/config");

const app = express();

mongoose
  .connect(
    `mongodb://${MONGO_USER}:${MONGO_PASSWORD}@${MONGO_IP}:${MONGO_PORT}/?authSource=admin`
  )
  .then(() => console.log("connected to db"))
  .catch((e) => console.log(e));

app.get("/", (req, res) => {
  res.send("<h2>Helllo</h2>");
});

const port = process.env.PORT || 3300;

app.listen(port, () => console.log(`listening on port ${port}`));

CREATE TABLE Users(
  UserID STRING(MAX) NOT NULL,
) PRIMARY KEY(UserID);

CREATE TABLE Blocks(
  UserID STRING(MAX) NOT NULL,
  BlockedUserID STRING(MAX) NOT NULL,
  CONSTRAINT FK_Blocks_UserID FOREIGN KEY(UserID) REFERENCES Users(UserID),
  CONSTRAINT FK_Blocks_BlockedUserID FOREIGN KEY(BlockedUserID) REFERENCES Users(UserID),
) PRIMARY KEY(UserID, BlockedUserID);

CREATE TABLE Comments(
  UserID STRING(MAX) NOT NULL,
  CommentID STRING(MAX) NOT NULL,
  SendingUserID STRING(MAX) NOT NULL,
  Content STRING(MAX) NOT NULL,
  CreateTime TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
  CONSTRAINT FK_Comments_UserID FOREIGN KEY(UserID) REFERENCES Users(UserID),
  CONSTRAINT FK_Comments_SendingUserID FOREIGN KEY(SendingUserID) REFERENCES Users(UserID),
) PRIMARY KEY(UserID, CommentID);
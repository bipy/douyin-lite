DROP TABLE IF EXISTS VIDEOS;
DROP TABLE IF EXISTS USERS;
DROP TABLE IF EXISTS FOLLOWS;
DROP TABLE IF EXISTS FAVORITES;
DROP TABLE IF EXISTS COMMENTS;

CREATE TABLE VIDEOS
(
    VIDEO_ID       INT PRIMARY KEY AUTO_INCREMENT COMMENT '视频ID',
    AUTHOR_ID      INT          NOT NULL COMMENT '视频发布者ID',
    PLAY_URL       VARCHAR(255) NOT NULL COMMENT '视频源地址',
    COVER_URL      VARCHAR(255) NOT NULL COMMENT '视频封面地址',
    TITLE          VARCHAR(32)  NOT NULL COMMENT '视频标题',
    FAVORITE_COUNT INT          NOT NULL DEFAULT 0 COMMENT '视频点赞数',
    COMMENT_COUNT  INT          NOT NULL DEFAULT 0 COMMENT '视频点赞数',
    CREATE_TIME    INT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP()) COMMENT '视频发布时间',
    INDEX VIDEOS_CREATE_TIME_index (CREATE_TIME) COMMENT '视频发布时间索引'
) CHARSET = utf8mb4 COMMENT '视频表';


CREATE TABLE USERS
(
    USER_ID        INT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    USER_NAME      VARCHAR(32) UNIQUE NOT NULL COMMENT '用户名',
    USER_PASSWORD  CHAR(60)           NOT NULL COMMENT '用户密码',
    FOLLOW_COUNT   INT                NOT NULL DEFAULT 0 COMMENT '关注数量',
    FOLLOWER_COUNT INT                NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    CREATE_TIME    INT UNSIGNED       NOT NULL DEFAULT (UNIX_TIMESTAMP()) COMMENT '用户创建时间',
    UNIQUE INDEX USER_NAME_INDEX (USER_NAME) COMMENT ''
) CHARSET = utf8mb4 COMMENT '用户表';

CREATE TABLE FOLLOWS
(
    A_ID        INT          NOT NULL COMMENT '关注',
    B_ID        INT          NOT NULL COMMENT '被关注',
    CREATE_TIME INT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP()) COMMENT '创建时间',
    PRIMARY KEY (A_ID, B_ID)
) COMMENT '关注表';

CREATE TABLE FAVORITES
(
    USER_ID     INT          NOT NULL COMMENT '用户ID',
    VIDEO_ID    INT          NOT NULL COMMENT '视频ID',
    CREATE_TIME INT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP()) COMMENT '创建时间',
    PRIMARY KEY (USER_ID, VIDEO_ID)
) COMMENT '点赞表';

CREATE TABLE COMMENTS
(
    COMMENT_ID   INT PRIMARY KEY AUTO_INCREMENT COMMENT '评论ID',
    USER_ID      INT          NOT NULL COMMENT '用户ID',
    VIDEO_ID     INT          NOT NULL COMMENT '视频ID',
    COMMENT_TEXT VARCHAR(255) NOT NULL COMMENT '评论内容',
    DELETED      TINYINT      NOT NULL DEFAULT 0 COMMENT '是否删除',
    CREATE_TIME  INT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP()) COMMENT '创建时间'
) COMMENT '评论表';

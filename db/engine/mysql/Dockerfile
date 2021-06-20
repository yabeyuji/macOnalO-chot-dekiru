FROM mysql:5.7.16
# RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
# ENV LANG ja_JP.UTF-:5.7.16

# 初期化用 sql ファイルを、コンテナ内の所定のディレクトリにコピー
# COPY ./init/*.sql /docker-entrypoint-initdb.d/

RUN rm -rf ./data/*

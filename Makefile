APPNAME="go-task"
OUTDIR="./dist"
DSN="mysql://$(MYSQL_DSN)"
MIGRATION_DIR="file://db/migration"


# テスト + ビルドを行います。
all: test build

# 一時ファイルの削除などを行います。
clean:
	@rm -rf $(OUTDIR)
	@mkdir $(OUTDIR)

# アプリケーションの実行に必要な依存関係をインストールします。
depend:
	@go mod tidy

# アプリケーションをテストします。
test: depend 
	@go test ./...

# アプリケーションをビルドします。
build: clean depend test
	@go build -o $(OUTDIR)/$(APPNAME)

# アプリケーションを実行します。
# freshパッケージにより、ファイルが保存される度に自動ビルドされます。
run: depend
	@go get github.com/pilu/fresh
	fresh
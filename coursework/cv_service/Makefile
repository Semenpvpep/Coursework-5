# Укажите путь к protoc
PROTOC = protoc

# Укажите путь к папке с .proto файлами
PROTO_DIR = ./proto

# Укажите путь к папке для сгенерированных файлов
OUT_DIR = ./api

# Укажите расширения, которые вы хотите сгенерировать (например, для Go, Python и т.д.)
# Добавьте нужные вам плагины
PROTOC_GEN_GO = --go_out=$(OUT_DIR) --go_opt=paths=source_relative
# Для Python
# PROTOC_GEN_PY = --python_out=$(OUT_DIR)

# Список .proto файлов для генерации
PROTO_FILES = $(PROTO_DIR)/your_proto_file.proto
# Цель по умолчанию
all: generate
protoc -I proto proto/cv.proto --go_out=./api/proto --go_opt=paths=source_relative --go-grpc_out=./api/proto --go-grpc_opt=paths=source_relative
# Цель для генерации .pb файлов
generate:
	$(PROTOC) $(PROTOC_GEN_GO) $(PROTO_FILES)
	# Для Python:
	# $(PROTOC) $(PROTOC_GEN_PY) $(PROTO_FILES)

# Очистка сгенерированных файлов
clean:
	rm -f $(OUT_DIR)/*.pb.go
	# Для Python:
	# rm -f $(OUT_DIR)/*.pb.py

.PHONY: all generate clean
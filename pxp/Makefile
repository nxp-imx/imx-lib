CC ?=$(CROSS_COMPILE)gcc
AR ?=$(CROSS_COMPILE)ar
CFLAGS ?=-O2

# list of platforms which want this test case
INCLUDE_LIST:=IMX50 IMX51 IMX5 IMX6Q IMX6S IMX7 IMX6UL IMX8ULP

OBJ = pxp_lib.o

LIBNAME = libpxp
SONAMEVERSION=0
LFLAGS += -lpthread

libdir ?= "/usr/lib"

ifeq ($(PLATFORM), $(findstring $(PLATFORM), $(INCLUDE_LIST)))

all: $(LIBNAME).so $(LIBNAME).a

install: install_headers
	@mkdir -p $(DEST_DIR)$(libdir)
	cp -P $(LIBNAME).* $(DEST_DIR)$(libdir)

install_headers:
	@mkdir -p $(DEST_DIR)/usr/include
	cp pxp_lib.h $(DEST_DIR)/usr/include
else
all install :
endif

%.o: %.c
	$(CC) -D$(PLATFORM) $(INCLUDE) -Wall -fPIC $(CFLAGS) -c $^ -o $@

$(LIBNAME).so.$(SONAMEVERSION): $(OBJ)
	$(CC) -shared -nostartfiles -Wl,-soname,$@ $(LFLAGS) $^ -o $@ $(LDFLAGS)

$(LIBNAME).so: $(LIBNAME).so.$(SONAMEVERSION)
	ln -s $< $@

$(LIBNAME).a: $(OBJ)
	$(AR) -rc $@  $^

.PHONY: clean
clean:
	rm -f $(LIBNAME).* $(OBJ)

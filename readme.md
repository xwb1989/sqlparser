[![Build Status](https://travis-ci.org/xwb1989/sqlparser.svg?branch=master)](https://travis-ci.org/xwb1989/sqlparser)

# Notice

The backbone of this repo is extracted from [youtube/vitess](https://github.com/youtube/vitess).

Inside youtube/vitess there is a very nicely written sql parser. However as it's not a self-contained application, I created this one. 
It applies the same LICENSE as youtube/vitess.

# Usage

```go
    import (
        "github.com/xwb1989/sqlparser"
    )
```

Then use

```go    
    sqlparser.Parse(sql)
```

See `parse_test.go` for more `Parse` usage.

# Porting Instructions

You only need the below if you plan to try and keep this library up to date with [youtube/vitess](https://github.com/youtube/vitess).

## Keeping up to date

```bash
cd $GOPATH/src/github.com/youtube/vitess/go

# Create patches for everything that changed
LASTIMPORT=5d7c5b2a2f6f1766843729a89b649e182441b529
git format-patch $LASTIMPORT vt/sqlparser
git format-patch $LASTIMPORT sqltypes
git format-patch $LASTIMPORT bytes2
git format-patch $LASTIMPORT hack

# Apply them to the repo
cd $GOPATH/src/github.com/xwb1989/sqlparser
git am -p4 ../../youtube/vitess/go/*.patch
```

## Fresh install

```bash
cd $GOPATH/src/github.com/xwb1989/sqlparser

# Copy all the code
cp -pr ../../youtube/vitess/go/vt/sqlparser/ .
cp -pr ../../youtube/vitess/go/sqltypes dependency
cp -pr ../../youtube/vitess/go/bytes2 dependency
cp -pr ../../youtube/vitess/go/hack dependency

# Delete some code we haven't ported
rm dependency/sqltypes/arithmetic.go dependency/sqltypes/arithmetic_test.go dependency/sqltypes/event_token.go dependency/sqltypes/event_token_test.go dependency/sqltypes/proto3.go dependency/sqltypes/proto3_test.go dependency/sqltypes/query_response.go dependency/sqltypes/result.go dependency/sqltypes/result_test.go

# Some automated fixes

# Fix imports
sed -i '.bak' 's_github.com/youtube/vitess/go/vt/proto/query_github.com/xwb1989/sqlparser/dependency/querypb_g' *.go dependency/sqltypes/*.go
sed -i '.bak' 's_github.com/youtube/vitess/go/_github.com/xwb1989/sqlparser/dependency/_g' *.go dependency/sqltypes/*.go

# Copy the proto, but basically drop everything we don't want
cp -pr ../../youtube/vitess/go/vt/proto/query dependency/querypb
sed -i '.bak' 's_.*Descriptor.*__g' dependency/querypb/*.go
sed -i '.bak' 's_.*ProtoMessage.*__g' dependency/querypb/*.go

sed -i '.bak' 's/proto.CompactTextString(m)/"TODO"/g' dependency/querypb/*.go
sed -i '.bak' 's/proto.EnumName/EnumName/g' dependency/querypb/*.go

sed -i '.bak' 's/proto.Equal/reflect.DeepEqual/g' dependency/sqltypes/*.go

# Remove the error library
sed -i '.bak' 's/vterrors.Errorf([^,]*, /fmt.Errorf(/g' *.go dependency/sqltypes/*.go
sed -i '.bak' 's/vterrors.New([^,]*, /errors.New(/g' *.go dependency/sqltypes/*.go
```

## Testing

```bash
cd $GOPATH/src/github.com/xwb1989/sqlparser

# Test, fix and repeat
go test ./...

# Finally make some diffs (for later reference)
cd $GOPATH/src/github.com
diff -u youtube/vitess/go/sqltypes/        xwb1989/sqlparser/dependency/sqltypes/ > xwb1989/sqlparser/patches/sqltypes.patch
diff -u youtube/vitess/go/bytes2/          xwb1989/sqlparser/dependency/bytes2/   > xwb1989/sqlparser/patches/bytes2.patch
diff -u youtube/vitess/go/vt/proto/query/  xwb1989/sqlparser/dependency/querypb/  > xwb1989/sqlparser/patches/querypb.patch
diff -u youtube/vitess/go/vt/sqlparser/    xwb1989/sqlparser/                     > xwb1989/sqlparser/patches/sqlparser.patch
```
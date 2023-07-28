table "employees" {
  schema = schema.public
  column "email" {
    null = false
    type = text
  }
  column "name" {
    null = true
    type = text
  }
  column "profile_picture" {
    null = true
    type = text
  }
  column "org_id" {
    null = true
    type = bigint
  }
  column "role" {
    null = true
    type = text
  }
  column "managed_by" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.email]
  }
}
table "organizations" {
  schema = schema.public
  column "id" {
    null    = false
    type    = bigint
    default = sql("unique_rowid()")
  }
  column "name" {
    null = true
    type = text
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "created_by" {
    null = true
    type = text
  }
  column "is_active" {
    null = true
    type = boolean
  }
  column "is_deleted" {
    null = true
    type = boolean
  }
  primary_key {
    columns = [column.id]
  }
}
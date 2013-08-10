package reflections_test

import (
    "fmt"
    "log"
    "github.com/oleiade/reflections"
)

type MyStruct struct {
    FirstField  string      `matched:"first tag"`
    SecondField int         `matched:"second tag"`
    ThirdField  string      `unmatched:"third tag"`
}

func ExampleGetField() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    fieldsToExtract := []string{"FirstField", "ThirdField"}

    for _, fieldName := range fieldsToExtract {
        value, err := reflections.GetField(s, fieldName)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(value)
    }
}

func ExampleHasField() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    // has == true
    has, _ := reflections.HasField(s, "FirstField")
    fmt.Println(has)

    // has == false
    has, _ = reflections.HasField(s, "FourthField")
    fmt.Println(has)
}

func ExampleFields() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    var fields []string

    // Fields will list every structure exportable fields.
    // Here, it's content would be equal to:
    // []string{"FirstField", "SecondField", "ThirdField"}
    fields, _ = reflections.Fields(s)
    fmt.Println(fields)
}

func ExampleItems() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    var structItems map[string]interface{}

    // Items will return a field name to
    // field value map
    structItems, _ = reflections.Items(s)
    fmt.Println(structItems)
}

func ExampleTags() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    var structTags map[string]string

    // Tags will return a field name to tag content
    // map. Nota that only field with the tag name
    // you've provided which will be matched.
    // Here structTags will contain:
    // {
    //     "FirstField": "first tag",
    //     "SecondField": "second tag",
    // }
    structTags, _ = reflections.Tags(s, "matched")
    fmt.Println(structTags)
}

func ExampleSetField() {
    s := MyStruct {
        FirstField: "first value",
        SecondField: 2,
        ThirdField: "third value",
    }

    // In order to be able to set the structure's values,
    // a pointer to it has to be passed to it.
    err := reflections.SetField(&s, "FirstField", "new value")
    if err != nil {
        log.Fatal(err)
    }

    // If you try to set a field's value using the wrong type,
    // an error will be returned
    err = reflections.SetField(&s, "FirstField", 123)  // err != nil
    if err != nil {
        log.Fatal(err)
    }
}
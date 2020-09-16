import json

def add_missing_descriptions(swagger):
    new_dict = swagger
    for path_key, path in swagger["paths"].items():
        for response_types_key, response_types in path.items():
            if "400" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["400"]["description"] = "Bad Request"
            if "401" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["401"]["description"] = "Unauthorised"
            if "403" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["403"]["description"] = "No Permission Response"
            if "404" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["404"]["description"] = "Not Found"
            if "409" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["409"]["description"] = "Request Conflict"
            if "415" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["415"]["description"] = "Unsupported Media Type"
            if "502" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["502"]["description"] = "Bad Gateway"
            if "504" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["504"]["description"] = "Gateway Timeout"
            if "200" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["200"]["description"] = "OK"
            if "202" in response_types["responses"]:
                new_dict["paths"][path_key][response_types_key]["responses"]["202"]["description"] = "Accepted"
    return new_dict


with open("raw.json") as raw_file:
    data = json.load(raw_file)

    # Iterate through dictionary, and return all items that match filter:
    keys = add_missing_descriptions(data)
    
    with open('swagger.json', 'w') as out_file:
        json.dump(data, out_file, sort_keys=True, indent=2)
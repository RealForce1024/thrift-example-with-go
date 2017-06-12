//namespace py example
//
//struct Data {
//  1: string text
//}
//
//service format_data {
//    Data format_data(1:Data data),
//}


namespace py example ##
namespace java example ##

struct Data {
    1: string text
}

service format_data {
    Data do_format(1:Data data),
}
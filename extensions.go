/*
 * getdns extensions.
 * Currently, we only support the boolean extensions, but this will
 * need to be extended to support extensions that support other types
 * of data such as "add_opt_parameters" and "specify_class".
 */

package getdns

type Extension map[string]bool

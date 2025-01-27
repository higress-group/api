// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using AuthorizationPolicy within kubernetes types, where deepcopy-gen is used.
func (in *AuthorizationPolicy) DeepCopyInto(out *AuthorizationPolicy) {
	p := proto.Clone(in).(*AuthorizationPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationPolicy. Required by controller-gen.
func (in *AuthorizationPolicy) DeepCopy() *AuthorizationPolicy {
	if in == nil {
		return nil
	}
	out := new(AuthorizationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationPolicy. Required by controller-gen.
func (in *AuthorizationPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AuthorizationPolicy_ExtensionProvider within kubernetes types, where deepcopy-gen is used.
func (in *AuthorizationPolicy_ExtensionProvider) DeepCopyInto(out *AuthorizationPolicy_ExtensionProvider) {
	p := proto.Clone(in).(*AuthorizationPolicy_ExtensionProvider)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationPolicy_ExtensionProvider. Required by controller-gen.
func (in *AuthorizationPolicy_ExtensionProvider) DeepCopy() *AuthorizationPolicy_ExtensionProvider {
	if in == nil {
		return nil
	}
	out := new(AuthorizationPolicy_ExtensionProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationPolicy_ExtensionProvider. Required by controller-gen.
func (in *AuthorizationPolicy_ExtensionProvider) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule within kubernetes types, where deepcopy-gen is used.
func (in *Rule) DeepCopyInto(out *Rule) {
	p := proto.Clone(in).(*Rule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule_From within kubernetes types, where deepcopy-gen is used.
func (in *Rule_From) DeepCopyInto(out *Rule_From) {
	p := proto.Clone(in).(*Rule_From)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule_From. Required by controller-gen.
func (in *Rule_From) DeepCopy() *Rule_From {
	if in == nil {
		return nil
	}
	out := new(Rule_From)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule_From. Required by controller-gen.
func (in *Rule_From) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule_To within kubernetes types, where deepcopy-gen is used.
func (in *Rule_To) DeepCopyInto(out *Rule_To) {
	p := proto.Clone(in).(*Rule_To)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule_To. Required by controller-gen.
func (in *Rule_To) DeepCopy() *Rule_To {
	if in == nil {
		return nil
	}
	out := new(Rule_To)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule_To. Required by controller-gen.
func (in *Rule_To) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Source within kubernetes types, where deepcopy-gen is used.
func (in *Source) DeepCopyInto(out *Source) {
	p := proto.Clone(in).(*Source)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source. Required by controller-gen.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Source. Required by controller-gen.
func (in *Source) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Operation within kubernetes types, where deepcopy-gen is used.
func (in *Operation) DeepCopyInto(out *Operation) {
	p := proto.Clone(in).(*Operation)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation. Required by controller-gen.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Operation. Required by controller-gen.
func (in *Operation) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Condition within kubernetes types, where deepcopy-gen is used.
func (in *Condition) DeepCopyInto(out *Condition) {
	p := proto.Clone(in).(*Condition)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using StringMatch within kubernetes types, where deepcopy-gen is used.
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	p := proto.Clone(in).(*StringMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

package tidy

/*
#cgo CFLAGS: -I/usr/include/tidy
#cgo LDFLAGS: -ltidy -L/usr/local/lib 
#include <tidy.h>
#include <buffio.h>
#include <errno.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

const False int = 0
const True int = 1
const Auto int = 2

// HTML, XHTML, XML Options

// This option specifies if Tidy should add the XML declaration when outputting XML or XHTML. Note that if the input already includes an <?xml ... ?> declaration then this option will be ignored. If the encoding for the output is different from "ascii", one of the utf encodings or "raw", the declaration is always added as required by the XML standard.
func (this *Tidy) AddXmlDecl(val bool) (bool, error) {
	return this.optSetBool(C.TidyXmlDecl, cBool(val))
}

// This option specifies if Tidy should add xml:space="preserve" to elements such as <PRE>, <STYLE> and <SCRIPT> when generating XML. This is needed if the whitespace in such elements is to be parsed appropriately without having access to the DTD.
func (this *Tidy) AddXmlSpace(val bool) (bool, error) {
	return this.optSetBool(C.TidyXmlSpace, cBool(val))
}

// This option specifies the default "alt=" text Tidy uses for <IMG> attributes. This feature is dangerous as it suppresses further accessibility warnings. You are responsible for making your documents accessible to people who can not see the images!
func (this *Tidy) AltText(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyXmlSpace, v)
}

// This option specifies if Tidy should change the parsing of processing instructions to require ?> as the terminator rather than >. This option is automatically set if the input is in XML.
func (this *Tidy) AssumeXmlProcins(val bool) (bool, error) {
	return this.optSetBool(C.TidyXmlPIs, cBool(val))
}

// This option specifies if Tidy should strip Microsoft specific HTML from Word 2000 documents, and output spaces rather than non-breaking spaces where they exist in the input.
func (this *Tidy) Bare(val bool) (bool, error) {
	return this.optSetBool(C.TidyMakeBare, cBool(val))
}

// This option specifies if Tidy should strip out surplus presentational tags and attributes replacing them by style rules and structural markup as appropriate. It works well on the HTML saved by Microsoft Office products.
func (this *Tidy) Clean(val bool) (bool, error) {
	return this.optSetBool(C.TidyMakeClean, cBool(val))
}

// This option specifies the prefix that Tidy uses for styles rules. By default, "c" will be used.
func (this *Tidy) CssPrefix(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyCSSPrefix, v)
}

// This option specifies if Tidy should decorate inferred UL elements with some CSS markup to avoid indentation to the right.
func (this *Tidy) DecorateInferredUl(val bool) (bool, error) {
	return this.optSetBool(C.TidyDecorateInferredUL, cBool(val))
}

// This option specifies the DOCTYPE declaration generated by Tidy. If set to "omit" the output won't contain a DOCTYPE declaration. If set to "auto" (the default) Tidy will use an educated guess based upon the contents of the document. If set to "strict", Tidy will set the DOCTYPE to the strict DTD. If set to "loose", the DOCTYPE is set to the loose (transitional) DTD. Alternatively, you can supply a string for the formal public identifier (FPI).
//
// For example: 
// doctype: "-//ACME//DTD HTML 3.14159//EN"
//
// If you specify the FPI for an XHTML document, Tidy will set the system identifier to an empty string. For an HTML document, Tidy adds a system identifier only if one was already present in order to preserve the processing mode of some browsers. Tidy leaves the DOCTYPE for generic XML documents unchanged. --doctype omit implies --numeric-entities yes. This option does not offer a validation of the document conformance.
func (this *Tidy) Doctype(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyDoctype, v)
}

// This option specifies if Tidy should discard empty paragraphs.
func (this *Tidy) DropEmptyParas(val bool) (bool, error) {
	return this.optSetBool(C.TidyDropEmptyParas, cBool(val))
}

// This option specifies if Tidy should discard <FONT> and <CENTER> tags without creating the corresponding style rules. This option can be set independently of the clean option.
//func (this *Tidy) DropFontTags(val bool) (bool, error) {
//	return this.optSetBool(C.TidyDropFontTags, cBool(val))
//}

// This option specifies if Tidy should strip out proprietary attributes, such as MS data binding attributes.
func (this *Tidy) DropProprietaryAttributes(val bool) (bool, error) {
	return this.optSetBool(C.TidyDropPropAttrs, cBool(val))
}

// This option specifies if Tidy should insert a <P> element to enclose any text it finds in any element that allows mixed content for HTML transitional but not HTML strict.
func (this *Tidy) EncloseBlockText(val bool) (bool, error) {
	return this.optSetBool(C.TidyEncloseBlockText, cBool(val))
}

// This option specifies if Tidy should enclose any text it finds in the body element within a <P> element. This is useful when you want to take existing HTML and use it with a style sheet.
func (this *Tidy) EncloseText(val bool) (bool, error) {
	return this.optSetBool(C.TidyEncloseBodyText, cBool(val))
}

// This option specifies if Tidy should convert <![CDATA[]]> sections to normal text.
func (this *Tidy) EscapeCdata(val bool) (bool, error) {
	return this.optSetBool(C.TidyEscapeCdata, cBool(val))
}

// This option specifies if Tidy should replace backslash characters "\" in URLs by forward slashes "/".
func (this *Tidy) FixBackslash(val bool) (bool, error) {
	return this.optSetBool(C.TidyFixBackslash, cBool(val))
}

// This option specifies if Tidy should replace unexpected hyphens with "=" characters when it comes across adjacent hyphens. The default is yes. This option is provided for users of Cold Fusion which uses the comment syntax: <!--- --->
func (this *Tidy) FixBadComments(val bool) (bool, error) {
	return this.optSetBool(C.TidyFixComments, cBool(val))
}

// This option specifies if Tidy should check attribute values that carry URIs for illegal characters and if such are found, escape them as HTML 4 recommends.
func (this *Tidy) FixUri(val bool) (bool, error) {
	return this.optSetBool(C.TidyFixUri, cBool(val))
}

// This option specifies if Tidy should print out comments.
func (this *Tidy) HideComments(val bool) (bool, error) {
	return this.optSetBool(C.TidyHideComments, cBool(val))
}

// This option specifies if Tidy should omit optional end-tags when generating the pretty printed markup. This option is ignored if you are outputting to XML.
//func (this *Tidy) HideEndtags(val bool) (bool, error) {
//	return this.optSetBool(C.TidyHideEndTags, cBool(val))
//}

// This option specifies if Tidy should indent <![CDATA[]]> sections.
func (this *Tidy) IndentCdata(val bool) (bool, error) {
	return this.optSetBool(C.TidyIndentCdata, cBool(val))
}

// This option specifies if Tidy should use the XML parser rather than the error correcting HTML parser.
func (this *Tidy) InputXml(val bool) (bool, error) {
	return this.optSetBool(C.TidyXmlTags, cBool(val))
}

// This option specifies if Tidy should combine class names to generate a single new class name, if multiple class assignments are detected on an element.
func (this *Tidy) JoinClasses(val bool) (bool, error) {
	return this.optSetBool(C.TidyJoinClasses, cBool(val))
}

// This option specifies if Tidy should combine styles to generate a single new style, if multiple style values are detected on an element.
func (this *Tidy) JoinStyles(val bool) (bool, error) {
	return this.optSetBool(C.TidyJoinStyles, cBool(val))
}

// This option specifies if Tidy should ensure that whitespace characters within attribute values are passed through unchanged.
func (this *Tidy) LiteralAttributes(val bool) (bool, error) {
	return this.optSetBool(C.TidyLiteralAttribs, cBool(val))
}

// This option specifies if Tidy should replace any occurrence of <I> by <EM> and any occurrence of <B> by <STRONG>. In both cases, the attributes are preserved unchanged. This option can be set independently of the clean and drop-font-tags options.
func (this *Tidy) LogicalEmphasis(val bool) (bool, error) {
	return this.optSetBool(C.TidyLogicalEmphasis, cBool(val))
}

// This option specifies if Tidy should convert the value of an attribute that takes a list of predefined values to lower case. This is required for XHTML documents.
func (this *Tidy) LowerLiterals(val bool) (bool, error) {
	return this.optSetBool(C.TidyLowerLiterals, cBool(val))
}

// Can be used to modify behavior of -c (--clean yes) option. This option specifies if Tidy should merge nested <div> such as "<div><div>...</div></div>". If set to "auto", the attributes of the inner <div> are moved to the outer one. As well, nested <div> with ID attributes are not merged. If set to "yes", the attributes of the inner <div> are discarded with the exception of "class" and "style".
func (this *Tidy) MergeDivs(val int) (bool, error) {
	return this.optSetAutoBool(C.TidyMergeDivs, (C.ulong)(val))
}

// This option specifies if Tidy should allow numeric character references.
func (this *Tidy) Ncr(val bool) (bool, error) {
	return this.optSetBool(C.TidyNCR, cBool(val))
}

// This option specifies new block-level tags. This option takes a space or comma separated list of tag names. Unless you declare new tags, Tidy will refuse to generate a tidied file if the input includes previously unknown tags. Note you can't change the content model for elements such as <TABLE>, <UL>, <OL> and <DL>. This option is ignored in XML mode.
func (this *Tidy) NewBlocklevelTags(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyBlockTags, v)
}

// This option specifies new empty inline tags. This option takes a space or comma separated list of tag names. Unless you declare new tags, Tidy will refuse to generate a tidied file if the input includes previously unknown tags. Remember to also declare empty tags as either inline or blocklevel. This option is ignored in XML mode.
func (this *Tidy) NewEmptyTags(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyEmptyTags, v)
}

// This option specifies new non-empty inline tags. This option takes a space or comma separated list of tag names. Unless you declare new tags, Tidy will refuse to generate a tidied file if the input includes previously unknown tags. This option is ignored in XML mode.
func (this *Tidy) NewInlineTags(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyInlineTags, v)
}

// This option specifies new tags that are to be processed in exactly the same way as HTML's <PRE> element. This option takes a space or comma separated list of tag names. Unless you declare new tags, Tidy will refuse to generate a tidied file if the input includes previously unknown tags. Note you can not as yet add new CDATA elements (similar to <SCRIPT>). This option is ignored in XML mode.
func (this *Tidy) NewPreTags(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyPreTags, v)
}

// This option specifies if Tidy should output entities other than the built-in HTML entities (&amp;, &lt;, &gt; and &quot;) in the numeric rather than the named entity form. Only entities compatible with the DOCTYPE declaration generated are used. Entities that can be represented in the output encoding are translated correspondingly.
func (this *Tidy) NumericEntities(val bool) (bool, error) {
	return this.optSetBool(C.TidyNumEntities, cBool(val))
}

// This option specifies if Tidy should generate pretty printed output, writing it as HTML.
func (this *Tidy) OutputHtml(val bool) (bool, error) {
	return this.optSetBool(C.TidyHtmlOut, cBool(val))
}

// This option specifies if Tidy should generate pretty printed output, writing it as extensible HTML. This option causes Tidy to set the DOCTYPE and default namespace as appropriate to XHTML. If a DOCTYPE or namespace is given they will checked for consistency with the content of the document. In the case of an inconsistency, the corrected values will appear in the output. For XHTML, entities can be written as named or numeric entities according to the setting of the "numeric-entities" option. The original case of tags and attributes will be preserved, regardless of other options.
func (this *Tidy) OutputXhtml(val bool) (bool, error) {
	return this.optSetBool(C.TidyXhtmlOut, cBool(val))
}

// This option specifies if Tidy should pretty print output, writing it as well-formed XML. Any entities not defined in XML 1.0 will be written as numeric entities to allow them to be parsed by a XML parser. The original case of tags and attributes will be preserved, regardless of other options.
func (this *Tidy) OutputXml(val bool) (bool, error) {
	return this.optSetBool(C.TidyXmlOut, cBool(val))
}

// This option specifies if Tidy should output unadorned & characters as &amp;.
func (this *Tidy) QuoteAmpersand(val bool) (bool, error) {
	return this.optSetBool(C.TidyQuoteAmpersand, cBool(val))
}

// This option specifies if Tidy should output " characters as &quot; as is preferred by some editing environments. The apostrophe character ' is written out as &#39; since many web browsers don't yet support &apos;.
func (this *Tidy) QuoteMarks(val bool) (bool, error) {
	return this.optSetBool(C.TidyQuoteMarks, cBool(val))
}

// This option specifies if Tidy should output non-breaking space characters as entities, rather than as the Unicode character value 160 (decimal).
func (this *Tidy) QuoteNbsp(val bool) (bool, error) {
	return this.optSetBool(C.TidyQuoteNbsp, cBool(val))
}

const KeepFirst int = 0
const KeepLast int = 1

// This option specifies if Tidy should keep the first or last attribute, if an attribute is repeated, e.g. has two align attributes.
func (this *Tidy) RepeatedAttributes(val int) (bool, error) {
	switch val {
	case KeepFirst, KeepLast:
		return this.optSetInt(C.TidyDuplicateAttrs, (C.ulong)(val))
	}
	return false, errors.New("Argument val int is out of range (0,1)")
}

// This option specifies if Tidy should replace numeric values in color attributes by HTML/XHTML color names where defined, e.g. replace "#ffffff" with "white".
func (this *Tidy) ReplaceColor(val bool) (bool, error) {
	return this.optSetBool(C.TidyReplaceColor, cBool(val))
}

// This option specifies if Tidy should print only the contents of the body tag as an HTML fragment. If set to "auto", this is performed only if the body tag has been inferred. Useful for incorporating existing whole pages as a portion of another page. This option has no effect if XML output is requested.
func (this *Tidy) ShowBodyOnly(val int) (bool, error) {
	return this.optSetAutoBool(C.TidyBodyOnly, (C.ulong)(val))
}

// This option specifies if Tidy should output attribute names in upper case. The default is no, which results in lower case attribute names, except for XML input, where the original case is preserved.
func (this *Tidy) UppercaseAttributes(val bool) (bool, error) {
	return this.optSetBool(C.TidyUpperCaseAttrs, cBool(val))
}

// This option specifies if Tidy should output tag names in upper case. The default is no, which results in lower case tag names, except for XML input, where the original case is preserved.
func (this *Tidy) UppercaseTags(val bool) (bool, error) {
	return this.optSetBool(C.TidyUpperCaseTags, cBool(val))
}

// This option specifies if Tidy should go to great pains to strip out all the surplus stuff Microsoft Word 2000 inserts when you save Word documents as "Web pages". Doesn't handle embedded images or VML. You should consider using Word's "Save As: Web Page, Filtered".
func (this *Tidy) Word2000(val bool) (bool, error) {
	return this.optSetBool(C.TidyWord2000, cBool(val))
}

// Diagnostics Options

const TidyClassic int = 0
const Priority1Checks int = 1
const Priority2Checks int = 2
const Priority3Checks int = 3

// This option specifies what level of accessibility checking, if any, that Tidy should do. Level 0 is equivalent to Tidy Classic's accessibility checking. For more information on Tidy's accessibility checking, visit the Adaptive Technology Resource Centre at the University of Toronto.
func (this *Tidy) AccessibilityCheck(val int) (bool, error) {
	switch val {
	case TidyClassic, Priority1Checks, Priority2Checks, Priority3Checks:
		return this.optSetInt(C.TidyAccessibilityCheckLevel, (C.ulong)(val))
	}
	return false, errors.New("Argument val int is out of range (0,1,2,3)")
}

// This option specifies the number Tidy uses to determine if further errors should be shown. If set to 0, then no errors are shown.
func (this *Tidy) ShowErrors(val int) (bool, error) {
	return this.optSetInt(C.TidyShowErrors, (C.ulong)(val))
}

// This option specifies if Tidy should suppress warnings. This can be useful when a few errors are hidden in a flurry of warnings.
func (this *Tidy) ShowWarnings(val bool) (bool, error) {
	return this.optSetBool(C.TidyShowWarnings, cBool(val))
}

// Pretty Print Options

// This option specifies if Tidy should output a line break before each <BR> element.
func (this *Tidy) BreakBeforeBr(val bool) (bool, error) {
	return this.optSetBool(C.TidyBreakBeforeBR, cBool(val))
}

// This option specifies if Tidy should indent block-level tags. If set to "auto", this option causes Tidy to decide whether or not to indent the content of tags such as TITLE, H1-H6, LI, TD, TD, or P depending on whether or not the content includes a block-level element. You are advised to avoid setting indent to yes as this can expose layout bugs in some browsers.
func (this *Tidy) Indent(val int) (bool, error) {
	return this.optSetAutoBool(C.TidyIndentContent, (C.ulong)(val))
}

// This option specifies if Tidy should begin each attribute on a new line.
func (this *Tidy) IndentAttributes(val bool) (bool, error) {
	return this.optSetBool(C.TidyIndentAttributes, cBool(val))
}

// This option specifies the number of spaces Tidy uses to indent content, when indentation is enabled.
func (this *Tidy) IndentSpaces(val int) (bool, error) {
	return this.optSetInt(C.TidyIndentSpaces, (C.ulong)(val))
}

// This option specifies if Tidy should generate a pretty printed version of the markup. Note that Tidy won't generate a pretty printed version if it finds significant errors (see force-output).
func (this *Tidy) Markup(val bool) (bool, error) {
	return this.optSetBool(C.TidyShowMarkup, cBool(val))
}

// This option specifies if Tidy should line wrap after some Unicode or Chinese punctuation characters.
func (this *Tidy) PunctuationWrap(val bool) (bool, error) {
	return this.optSetBool(C.TidyPunctWrap, cBool(val))
}

const None int = 0
const Alpha int = 1

// Currently not used. Tidy Classic only.
//func (this *Tidy) Split(val bool) (bool, error) {
//	return this.optSetBool(C.TidyBurstSlides, cBool(val))
//}

// This option specifies the number of columns that Tidy uses between successive tab stops. It is used to map tabs to spaces when reading the input. Tidy never outputs tabs.
func (this *Tidy) TabSize(val int) (bool, error) {
	return this.optSetInt(C.TidyTabSize, (C.ulong)(val))
}

// This option specifies if Tidy should add some empty lines for readability.
func (this *Tidy) VerticalSpace(val bool) (bool, error) {
	return this.optSetBool(C.TidyVertSpace, cBool(val))
}

// This option specifies the right margin Tidy uses for line wrapping. Tidy tries to wrap lines so that they do not exceed this length. Set wrap to zero if you want to disable line wrapping.
func (this *Tidy) Wrap(val int) (bool, error) {
	return this.optSetInt(C.TidyWrapLen, (C.ulong)(val))
}

// This option specifies if Tidy should line wrap text contained within ASP pseudo elements, which look like: <% ... %>.
func (this *Tidy) WrapAsp(val bool) (bool, error) {
	return this.optSetBool(C.TidyWrapAsp, cBool(val))
}

// This option specifies if Tidy should line wrap attribute values, for easier editing. This option can be set independently of wrap-script-literals.
func (this *Tidy) WrapAttributes(val bool) (bool, error) {
	return this.optSetBool(C.TidyWrapAttVals, cBool(val))
}

// This option specifies if Tidy should line wrap text contained within JSTE pseudo elements, which look like: <# ... #>.
func (this *Tidy) WrapJste(val bool) (bool, error) {
	return this.optSetBool(C.TidyWrapJste, cBool(val))
}

// This option specifies if Tidy should line wrap text contained within PHP pseudo elements, which look like: <?php ... ?>.
func (this *Tidy) WrapPhp(val bool) (bool, error) {
	return this.optSetBool(C.TidyWrapPhp, cBool(val))
}

// This option specifies if Tidy should line wrap string literals that appear in script attributes. Tidy wraps long script string literals by inserting a backslash character before the line break.
func (this *Tidy) WrapScriptLiterals(val bool) (bool, error) {
	return this.optSetBool(C.TidyWrapScriptlets, cBool(val))
}

// This option specifies if Tidy should line wrap text contained within <![ ... ]> section tags.
func (this *Tidy) WrapSections(val bool) (bool, error) {
	return this.optSetBool(C.TidyWrapSection, cBool(val))
}

// Character Encoding Options

// Can be used to modify behavior of -c (--clean yes) option. If set to "yes" when using -c, &emdash;, &rdquo;, and other named character entities are downgraded to their closest ascii equivalents.
func (this *Tidy) AsciiChars(val bool) (bool, error) {
	return this.optSetBool(C.TidyAsciiChars, cBool(val))
}

const Raw int = 0
const Ascii int = 1
const Latin0 int = 2
const Latin1 int = 3
const Utf8 int = 4
const Iso2022 int = 5
const Mac int = 6
const Win1252 int = 7
const Ibm858 int = 8
const Utf16le int = 9
const Utf16be int = 10
const Utf16 int = 11
const Big5 int = 12
const Shiftjis int = 13

// This option specifies the character encoding Tidy uses for both the input and output. For ascii, Tidy will accept Latin-1 (ISO-8859-1) character values, but will use entities for all characters whose value > 127. For raw, Tidy will output values above 127 without translating them into entities. For latin1, characters above 255 will be written as entities. For utf8, Tidy assumes that both input and output is encoded as UTF-8. You can use iso2022 for files encoded using the ISO-2022 family of encodings e.g. ISO-2022-JP. For mac and win1252, Tidy will accept vendor specific character values, but will use entities for all characters whose value > 127. For unsupported encodings, use an external utility to convert to and from UTF-8.
func (this *Tidy) CharEncoding(val int) (bool, error) {
	switch val {
	case Raw, Ascii, Latin0, Latin1, Utf8, Iso2022, Mac, Win1252, Ibm858, Utf16le, Utf16be, Utf16, Big5, Shiftjis:
		return this.optSetInt(C.TidyCharEncoding, (C.ulong)(val))
	}
	return false, errors.New("Argument val int is out of range (0-13)")
}

// This option specifies the character encoding Tidy uses for the input. See char-encoding for more info.
func (this *Tidy) InputEncoding(val int) (bool, error) {
	switch val {
	case Raw, Ascii, Latin0, Latin1, Utf8, Iso2022, Mac, Win1252, Ibm858, Utf16le, Utf16be, Utf16, Big5, Shiftjis:
		return this.optSetInt(C.TidyInCharEncoding, (C.ulong)(val))
	}
	return false, errors.New("Argument val int is out of range (0-13)")
}

// Currently not used, but this option specifies the language Tidy uses (for instance "en").
//func (this *Tidy) Language(val string) (bool, error) {
//	v := (*C.tmbchar)(C.CString(val))
//	defer C.free(unsafe.Pointer(v))
//	return this.optSetString(C.TidyLanguage, v)
//}

const LF int = 0
const CRLF int = 1
const CR int = 2

// The default is appropriate to the current platform: CRLF on PC-DOS, MS-Windows and OS/2, CR on Classic Mac OS, and LF everywhere else (Unix and Linux).
func (this *Tidy) Newline(val int) (bool, error) {
	switch val {
	case LF, CRLF, CR:
		return this.optSetInt(C.TidyNewline, (C.ulong)(val))
	}
	return false, errors.New("Argument val int is out of range (0,1,2)")
}

// This option specifies if Tidy should write a Unicode Byte Order Mark character (BOM; also known as Zero Width No-Break Space; has value of U+FEFF) to the beginning of the output; only for UTF-8 and UTF-16 output encodings. If set to "auto", this option causes Tidy to write a BOM to the output only if a BOM was present at the beginning of the input. A BOM is always written for XML/XHTML output using UTF-16 output encodings.
func (this *Tidy) OutputBom(val int) (bool, error) {
	return this.optSetAutoBool(C.TidyOutputBOM, (C.ulong)(val))
}

// This option specifies the character encoding Tidy uses for the output. See char-encoding for more info. May only be different from input-encoding for Latin encodings (ascii, latin0, latin1, mac, win1252, ibm858).
func (this *Tidy) OutputEncoding(val int) (bool, error) {
	switch val {
	case Raw, Ascii, Latin0, Latin1, Utf8, Iso2022, Mac, Win1252, Ibm858, Utf16le, Utf16be, Utf16, Big5, Shiftjis:
		return this.optSetInt(C.TidyOutCharEncoding, (C.ulong)(val))
	}
	return false, errors.New("Argument val int is out of range (0-13)")
}

// Miscellaneous Options

// This option specifies the error file Tidy uses for errors and warnings. Normally errors and warnings are output to "stderr".
func (this *Tidy) ErrorFile(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyErrFile, v)
}

// This option specifies if Tidy should produce output even if errors are encountered. Use this option with care - if Tidy reports an error, this means Tidy was not able to, or is not sure how to, fix the error, so the resulting output may not reflect your intention.
func (this *Tidy) ForceOutput(val bool) (bool, error) {
	return this.optSetBool(C.TidyForceOutput, cBool(val))
}

// This option specifies if Tidy should change the format for reporting errors and warnings to a format that is more easily parsed by GNU Emacs.
func (this *Tidy) GnuEmacs(val bool) (bool, error) {
	return this.optSetBool(C.TidyEmacs, cBool(val))
}

// Used internally.
func (this *Tidy) GnuEmacsFile(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyEmacsFile, v)
}

// This option specifies if Tidy should keep the original modification time of files that Tidy modifies in place. The default is no. Setting the option to yes allows you to tidy files without causing these files to be uploaded to a web server when using a tool such as SiteCopy. Note this feature is not supported on some platforms.
func (this *Tidy) KeepTime(val bool) (bool, error) {
	return this.optSetBool(C.TidyKeepFileTimes, cBool(val))
}

// This option specifies the output file Tidy uses for markup. Normally markup is written to "stdout".
func (this *Tidy) OutputFile(val string) (bool, error) {
	v := (*C.tmbchar)(C.CString(val))
	defer C.free(unsafe.Pointer(v))
	return this.optSetString(C.TidyOutFile, v)
}

// This option specifies if Tidy should output the summary of the numbers of errors and warnings, or the welcome or informational messages.
func (this *Tidy) Quiet(val bool) (bool, error) {
	return this.optSetBool(C.TidyQuiet, cBool(val))
}

// Currently not used. Tidy Classic only.
//func (this *Tidy) SlideStyle(val string) (bool, error) {
//	v := (*C.tmbchar)(C.CString(val))
//	defer C.free(unsafe.Pointer(v))
//	return this.optSetString(C.TidySlideStyle, v)
//}

// This option specifies if Tidy should add a meta element to the document head to indicate that the document has been tidied. Tidy won't add a meta element if one is already present.
func (this *Tidy) TidyMark(val bool) (bool, error) {
	return this.optSetBool(C.TidyMark, cBool(val))
}

// This option specifies if Tidy should write back the tidied markup to the same file it read from. You are advised to keep copies of important files before tidying them, as on rare occasions the result may not be what you expect.
func (this *Tidy) WriteBack(val bool) (bool, error) {
	return this.optSetBool(C.TidyWriteBack, cBool(val))
}

const Autobool_false C.ulong = 0
const Autobool_true C.ulong = 1
const Autobool_auto C.ulong = 2

func (this *Tidy) optSetAutoBool(opt C.TidyOptionId, val C.ulong) (bool, error) {
	switch val {
	case Autobool_false, Autobool_true, Autobool_auto:
		return this.optSetInt(opt, val)
	}
	return false, errors.New("Argument val int is out of range (0,1,2)")
}

func (this *Tidy) optSetString(opt C.TidyOptionId, val *C.tmbchar) (bool, error) {
	if C.tidyOptSetValue(this.tdoc, opt, val) == 1 {
		return false, nil
	}
	return true, nil
}

func (this *Tidy) optSetInt(opt C.TidyOptionId, val C.ulong) (bool, error) {
	if C.tidyOptSetInt(this.tdoc, opt, val) == 1 {
		return false, nil
	}
	return true, nil
}

func (this *Tidy) optSetBool(opt C.TidyOptionId, val C.Bool) (bool, error) {
	var rc C.int = -1
	if C.tidyOptSetBool(this.tdoc, opt, val) == 1 {
		rc = C.tidySetErrorBuffer(this.tdoc, &this.errbuf) // Capture diagnostics
		if rc != 0 {
			return false, errors.New(C.GoStringN((*C.char)(unsafe.Pointer(this.errbuf.bp)), C.int(this.errbuf.size)))
		}
	}
	return true, nil
}

func cBool(val bool) C.Bool {
	var v uint32 = 0
	if val {
		v = 1
	}
	return C.Bool(v)
}

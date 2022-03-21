@@ -332,11 +332,11 @@ func (p *Peer) handle(msg Msg) error {
		msg.Discard()
		go SendItems(p.rw, pongMsg)
	case msg.Code == discMsg:
		var reason [1]DiscReason
		// This is the last message. We don't need to discard or
		// check errors because, the connection will be closed after it.
		rlp.Decode(msg.Payload, &reason)
		return reason[0]
		var m struct{ R DiscReason }
		rlp.Decode(msg.Payload, &m)
		return m.R
	case msg.Code < baseProtocolLength:
		// ignore other base protocol messages
		return msg.Discard()

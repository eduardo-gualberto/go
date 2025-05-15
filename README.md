# go

ideia:
    2 bots 
        a. bot do paciente
            - faz agendamentos
            - faz reagendamentos
            - faz cancelamentos
            - tira dúvidas
            - envia lembretes de agendamentos
            - pede feedback
        b. bot do profissional
            - faz agendamentos
            - faz reagendamentos
            - faz cancelamentos
            - faz relatório de agendamentos do dia


todo:
    - MessageReader deve retornar uma string
        - TextMessageReader lê o campo ...Text.Body
        - AudioMessageReader transcreve o áudio e retorna o resultado
        - (Image/Document)MessageReader extrair texto da image e associar a última mensagem não-imagética (???) ou apenas encaminhar para o profissional e envia uma mensagem para o usuário avisando

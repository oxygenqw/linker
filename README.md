# linker

## Tunnel
lt --port 8080 --subdomain linker

## Postgres
psql -h localhost -p 5432 -U linker -d linker_db

<!-- func NewPagesHandler(service *service.Service, logger *logger.Logger) (*PagesHandler, error) {
    tmpl, err := template.ParseGlob("./ui/pages/*.html")
    if err != nil {
        return nil, fmt.Errorf("failed to parse templates: %w", err)
    }
    
    return &PagesHandler{
        logger:    logger,
        service:   service,
        templates: tmpl,
    }, nil
} -->